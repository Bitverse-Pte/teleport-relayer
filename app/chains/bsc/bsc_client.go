package bsc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	xibcbsc "github.com/teleport-network/teleport/x/xibc/clients/light-clients/bsc/types"
	xibctendermint "github.com/teleport-network/teleport/x/xibc/clients/light-clients/tendermint/types"
	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"
	"github.com/teleport-network/teleport/x/xibc/core/host"
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"
	"github.com/teleport-network/teleport/x/xibc/exported"

	"github.com/teleport-network/teleport-relayer/app/chains/bsc/contracts"
	"github.com/teleport-network/teleport-relayer/app/chains/bsc/contracts/transfer"
	"github.com/teleport-network/teleport-relayer/app/interfaces"
	"github.com/teleport-network/teleport-relayer/app/types"
	"github.com/teleport-network/teleport-relayer/app/types/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var _ interfaces.IChain = new(Bsc)

const CtxTimeout = 100 * time.Second
const TryGetGasPriceTimeInterval = 10 * time.Second
const RetryTimeout = 15 * time.Second
const RetryTimes = 5

type Bsc struct {
	// uri                   string
	chainName             string
	chainType             string
	updateClientFrequency uint64
	contractCfgGroup      *ContractCfgGroup
	contracts             *contractGroup
	bindOpts              *bindOpts
	queryFilter           string
	slot                  int64
	maxGasPrice           *big.Int
	tipCoefficient        float64
	ethClient             *ethclient.Client
	gethCli               *gethclient.Client
	gethRpcCli            *rpc.Client
	l                     *sync.Mutex
}

func NewBsc(config *ChainConfig) (interfaces.IChain, error) {
	return newBsc(config)
}

func newBsc(config *ChainConfig) (*Bsc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	rpcClient, err := rpc.DialContext(ctx, config.ChainURI)
	if err != nil {
		return nil, err
	}

	ethClient := ethclient.NewClient(rpcClient)
	gethCli := gethclient.New(rpcClient)

	contractGroup, err := newContractGroup(ethClient, config.ContractCfgGroup)
	if err != nil {
		return nil, err
	}
	tmpBindOpts, err := newBindOpts(config.ContractBindOptsCfg)
	if err != nil {
		return nil, err
	}

	return &Bsc{
		chainType:             config.ChainType,
		chainName:             config.ChainName,
		updateClientFrequency: config.UpdateClientFrequency,
		contractCfgGroup:      config.ContractCfgGroup,
		ethClient:             ethClient,
		gethCli:               gethCli,
		gethRpcCli:            rpcClient,
		contracts:             contractGroup,
		bindOpts:              tmpBindOpts,
		queryFilter:           config.QueryFilter,
		slot:                  config.Slot,
		tipCoefficient:        config.TipCoefficient,
		maxGasPrice:           new(big.Int).SetUint64(config.ContractBindOptsCfg.MaxGasPrice),
		l:                     new(sync.Mutex),
	}, nil
}

func (b *Bsc) ClientUpdateValidate(revisionHeight, delayHeight, updateHeight uint64) (uint64, error) {

	return updateHeight, nil
}

func (b *Bsc) GetPackets(fromBlock, toBlock uint64, destChainType string) (*types.Packets, error) {
	bizPackets, err := b.getPackets(fromBlock, toBlock, "")
	if err != nil {
		return nil, err
	}
	ackPackets, err := b.getAckPackets(fromBlock, toBlock, "")
	if err != nil {
		return nil, err
	}

	packets := &types.Packets{
		BizPackets: bizPackets,
		AckPackets: ackPackets,
	}

	return packets, nil
}

func (b *Bsc) GetPacketsByHash(hash string) (*types.Packets, error) {
	transaction, err := b.TransactionByHash(hash)
	if err != nil {
		return nil, err
	}
	block, err := b.ethClient.BlockByHash(context.Background(), *transaction.BlockHash)
	if err != nil {
		return nil, err
	}
	height := block.NumberU64()
	bizPackets, err := b.getPackets(height, height, hash)
	if err != nil {
		return nil, err
	}
	ackPackets, err := b.getAckPackets(height, height, hash)
	if err != nil {
		return nil, err
	}
	packets := &types.Packets{
		BizPackets: bizPackets,
		AckPackets: ackPackets,
	}
	return packets, nil
}

func (b *Bsc) TransactionByHash(hash string) (*types.RpcTransaction, error) {
	var transaction *types.RpcTransaction
	err := b.gethRpcCli.CallContext(context.Background(), &transaction, "eth_getTransactionByHash", hash)
	return transaction, err
}

func (b *Bsc) GetProof(sourChainName, destChainName string, sequence uint64, height uint64, typ string) ([]byte, error) {
	pkConstr := xibcbsc.NewProofKeyConstructor(sourChainName, destChainName, sequence)
	var key []byte
	switch typ {
	case types.CommitmentPoof:
		key = pkConstr.GetPacketCommitmentProofKey()
	case types.AckProof:
		key = pkConstr.GetAckProofKey()
	default:
		return nil, errors.ErrGetProof
	}
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	address := common.HexToAddress(b.contractCfgGroup.Packet.Addr)
	result, err := b.getProof(ctx, address, []string{hexutil.Encode(key)}, new(big.Int).SetUint64(height))
	if err != nil {
		return nil, err
	}

	var storageProof []*xibcbsc.StorageResult
	for _, sp := range result.StorageProof {
		tmpStorageProof := &xibcbsc.StorageResult{
			Key:   sp.Key,
			Value: hexutil.EncodeBig(sp.Value),
			Proof: sp.Proof,
		}
		storageProof = append(storageProof, tmpStorageProof)
	}
	nonce := hexutil.EncodeUint64(result.Nonce)
	balance := hexutil.EncodeBig(result.Balance)
	proof := &xibcbsc.Proof{
		Address:      result.Address.String(),
		Balance:      balance,
		CodeHash:     result.CodeHash.String(),
		Nonce:        nonce,
		StorageHash:  result.StorageHash.String(),
		AccountProof: result.AccountProof,
		StorageProof: storageProof,
	}

	return json.Marshal(proof)
}

func (b *Bsc) RelayPackets(msgs sdk.Msg) (string, error) {
	b.l.Lock()
	defer b.l.Unlock()

	resultTx := &types.ResultTx{}
	switch msg := msgs.(type) {
	case *packettypes.MsgRecvPacket:
		tmpPack := contracts.PacketTypesPacket{
			Sequence:    msg.Packet.Sequence,
			Ports:       msg.Packet.Ports,
			DestChain:   msg.Packet.DestinationChain,
			SourceChain: msg.Packet.SourceChain,
			RelayChain:  msg.Packet.RelayChain,
			DataList:    msg.Packet.DataList,
		}
		height := contracts.HeightData{
			RevisionNumber: msg.ProofHeight.RevisionNumber,
			RevisionHeight: msg.ProofHeight.RevisionHeight,
		}

		if err := b.setPacketOpts(); err != nil {
			return "", err
		}
		result, err := b.contracts.Packet.RecvPacket(
			b.bindOpts.packetTransactOpts,
			tmpPack,
			msg.ProofCommitment,
			height,
		)
		if err != nil {
			return "", err
		}
		resultTx.Hash = result.Hash().String()

	case *packettypes.MsgAcknowledgement:
		tmpPack := contracts.PacketTypesPacket{
			Sequence:    msg.Packet.Sequence,
			Ports:       msg.Packet.Ports,
			DestChain:   msg.Packet.DestinationChain,
			SourceChain: msg.Packet.SourceChain,
			RelayChain:  msg.Packet.RelayChain,
			DataList:    msg.Packet.DataList,
		}
		height := contracts.HeightData{
			RevisionNumber: msg.ProofHeight.RevisionNumber,
			RevisionHeight: msg.ProofHeight.RevisionHeight,
		}

		if err := b.setPacketOpts(); err != nil {
			return "", err
		}

		result, err := b.contracts.Packet.AcknowledgePacket(
			b.bindOpts.packetTransactOpts,
			tmpPack, msg.Acknowledgement, msg.ProofAcked,
			height,
		)
		if err != nil {
			return "", err
		}
		resultTx.Hash = result.Hash().String()
	}
	if err := b.reTryEthResult(resultTx.Hash, 0); err != nil {
		return resultTx.Hash, err
	}
	return resultTx.Hash, nil
}

func (b *Bsc) GetCommitmentsPacket(sourChainName, destChainName string, sequence uint64) error {
	hashBytes, err := b.contracts.Packet.Commitments(nil, host.PacketCommitmentKey(sourChainName, destChainName, sequence))
	if err != nil {
		return err
	}
	expectByte := make([]byte, 32)
	if bytes.Equal(expectByte, hashBytes[:]) {
		return fmt.Errorf("commitment does not exist")
	}
	return nil
}

func (b *Bsc) GetReceiptPacket(sourChainName, destChainName string, sequence uint64) (bool, error) {
	return b.contracts.Packet.Receipts(nil, host.PacketReceiptKey(sourChainName, destChainName, sequence))
}

func (b *Bsc) GetBlockHeader(req *types.GetBlockHeaderReq) (exported.Header, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	blockRes, err := b.ethClient.BlockByNumber(ctx, new(big.Int).SetUint64(req.LatestHeight))
	if err != nil {
		return nil, err
	}
	header := blockRes.Header()
	bscHeader := &xibcbsc.BscHeader{
		ParentHash:  header.ParentHash,
		UncleHash:   header.UncleHash,
		Coinbase:    header.Coinbase,
		Root:        header.Root,
		TxHash:      header.TxHash,
		ReceiptHash: header.ReceiptHash,
		Bloom:       xibcbsc.Bloom(header.Bloom),
		Difficulty:  header.Difficulty,
		Number:      header.Number,
		GasLimit:    header.GasLimit,
		GasUsed:     header.GasUsed,
		Time:        header.Time,
		Extra:       header.Extra,
		MixDigest:   header.MixDigest,
		Nonce:       xibcbsc.BlockNonce(header.Nonce),
	}
	protoHeader := bscHeader.ToHeader()
	return &protoHeader, nil
}

func (b *Bsc) GetBlockTimestamp(height uint64) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	blockRes, err := b.ethClient.BlockByNumber(ctx, new(big.Int).SetUint64(height))
	if err != nil {
		return 0, err
	}
	return blockRes.Time(), nil
}

func (b *Bsc) GetLightClientState(chainName string) (exported.ClientState, error) {
	latestHeight, err := b.contracts.Client.GetLatestHeight(nil, chainName)
	if err != nil {
		return nil, err
	}
	return &xibctendermint.ClientState{
		LatestHeight: clienttypes.Height{
			RevisionHeight: latestHeight.RevisionHeight,
			RevisionNumber: latestHeight.RevisionNumber,
		},
	}, nil
}

func (b *Bsc) GetLightClientConsensusState(chainName string, Height uint64) (exported.ConsensusState, error) {
	return nil, nil
}

func (b *Bsc) GetLatestHeight() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	return b.ethClient.BlockNumber(ctx)
}

func (b *Bsc) GetLightClientDelayHeight(s string) (uint64, error) {
	return 0, nil
}

func (b *Bsc) GetLightClientDelayTime(s string) (uint64, error) {
	return 0, nil
}

func (b *Bsc) UpdateClient(header exported.Header, chainName string) error {
	b.l.Lock()
	defer b.l.Unlock()

	h, ok := header.(*xibctendermint.Header)
	if !ok {
		return fmt.Errorf("invalid header type")
	}
	headerBytes, _ := h.Marshal()
	if err := b.setClientOpts(); err != nil {
		return err
	}
	result, err := b.contracts.Client.UpdateClient(b.bindOpts.client, chainName, headerBytes)
	if err != nil {
		return err
	}
	if err := b.reTryEthResult(result.Hash().String(), 0); err != nil {
		return err // TODO: warp
	}
	return nil
}

func (b *Bsc) BatchUpdateClient(headers []exported.Header, chainName string) error {
	return nil
}

func (b *Bsc) GetResult(hash string) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()

	cmnHash := common.HexToHash(hash)
	result, err := b.ethClient.TransactionReceipt(ctx, cmnHash)
	if err != nil {
		return 0, err
	}
	return result.Status, nil
}

func (b *Bsc) ChainName() string {
	return b.chainName
}

func (b *Bsc) UpdateClientFrequency() uint64 {
	return b.updateClientFrequency
}

func (b *Bsc) ChainType() string {
	return types.BSC
}

func (b *Bsc) getProof(ctx context.Context, account common.Address, keys []string, blockNumber *big.Int) (*gethclient.AccountResult, error) {
	type storageResult struct {
		Key   string       `json:"key"`
		Value *hexutil.Big `json:"value"`
		Proof []string     `json:"proof"`
	}

	type accountResult struct {
		Address      common.Address  `json:"address"`
		AccountProof []string        `json:"accountProof"`
		Balance      *hexutil.Big    `json:"balance"`
		CodeHash     common.Hash     `json:"codeHash"`
		Nonce        hexutil.Uint64  `json:"nonce"`
		StorageHash  common.Hash     `json:"storageHash"`
		StorageProof []storageResult `json:"storageProof"`
	}

	var res accountResult
	if err := b.gethRpcCli.CallContext(ctx, &res, "eth_getProof", account, keys, toBlockNumArg(blockNumber)); err != nil {
		return nil, err
	}

	// Turn hexutils back to normal datatypes
	storageResults := make([]gethclient.StorageResult, 0, len(res.StorageProof))
	for _, st := range res.StorageProof {
		storageResults = append(storageResults, gethclient.StorageResult{
			Key:   st.Key,
			Value: st.Value.ToInt(),
			Proof: st.Proof,
		})
	}
	result := &gethclient.AccountResult{
		Address:      res.Address,
		AccountProof: res.AccountProof,
		Balance:      res.Balance.ToInt(),
		Nonce:        uint64(res.Nonce),
		CodeHash:     res.CodeHash,
		StorageHash:  res.StorageHash,
		StorageProof: storageResults,
	}

	return result, nil
}

func (b *Bsc) reTryEthResult(hash string, n uint64) error {
	if n == RetryTimes {
		return fmt.Errorf("retry %d times and return error", RetryTimes)
	}
	txStatus, err := b.GetResult(hash)
	if err != nil {
		time.Sleep(RetryTimeout)
		return b.reTryEthResult(hash, n+1)
	}
	if txStatus == 0 {
		return fmt.Errorf("txStatus == 0,tx failed")
	}
	return nil
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Cmp(big.NewInt(-1)) == 0 {
		return "pending"
	}
	return hexutil.EncodeBig(number)
}

// get packets from block
func (b *Bsc) getPackets(fromBlock, toBlock uint64, hash string) ([]packettypes.Packet, error) {
	if strings.Contains(b.queryFilter, types.Packet) {
		return nil, nil
	}
	address := common.HexToAddress(b.contractCfgGroup.Packet.Addr)
	topic := b.contractCfgGroup.Packet.Topic
	logs, err := b.getLogs(address, topic, fromBlock, toBlock)
	if err != nil {
		return nil, err
	}

	var bizPackets []packettypes.Packet
	for _, log := range logs {
		if hash != "" && !bytes.Equal(log.TxHash.Bytes(), common.HexToHash(hash).Bytes()) {
			continue
		}
		packSent, err := b.contracts.Packet.ParsePacketSent(log)
		if err != nil {
			return nil, err
		}
		tmpPack := packettypes.Packet{
			Sequence:         packSent.Packet.Sequence,
			DataList:         packSent.Packet.DataList,
			SourceChain:      packSent.Packet.SourceChain,
			DestinationChain: packSent.Packet.DestChain,
			Ports:            packSent.Packet.Ports,
			RelayChain:       packSent.Packet.RelayChain,
		}
		bizPackets = append(bizPackets, tmpPack)
	}
	return bizPackets, nil
}

// get ack packets from block
func (b *Bsc) getAckPackets(fromBlock, toBlock uint64, hash string) ([]types.AckPacket, error) {
	if strings.Contains(b.queryFilter, types.Ack) {
		return nil, nil
	}
	address := common.HexToAddress(b.contractCfgGroup.AckPacket.Addr)
	topic := b.contractCfgGroup.AckPacket.Topic
	logs, err := b.getLogs(address, topic, fromBlock, toBlock)
	if err != nil {
		return nil, err
	}

	var ackPackets []types.AckPacket
	for _, log := range logs {
		if hash != "" && !bytes.Equal(log.TxHash.Bytes(), common.HexToHash(hash).Bytes()) {
			continue
		}
		ackWritten, err := b.contracts.Packet.ParseAckWritten(log)
		if err != nil {
			return nil, err
		}
		tmpAckPack := types.AckPacket{}
		tmpAckPack.Packet = packettypes.Packet{
			Sequence:         ackWritten.Packet.Sequence,
			DataList:         ackWritten.Packet.DataList,
			SourceChain:      ackWritten.Packet.SourceChain,
			DestinationChain: ackWritten.Packet.DestChain,
			Ports:            ackWritten.Packet.Ports,
			RelayChain:       ackWritten.Packet.RelayChain,
		}
		tmpAckPack.Acknowledgement = ackWritten.Ack
		ackPackets = append(ackPackets, tmpAckPack)
	}
	return ackPackets, nil
}

func (b *Bsc) getLogs(address common.Address, topic string, fromBlock, toBlock uint64) ([]ethtypes.Log, error) {
	filter := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []common.Address{address},
		Topics:    [][]common.Hash{{ethcrypto.Keccak256Hash([]byte(topic))}},
	}
	return b.ethClient.FilterLogs(context.Background(), filter)
}

func (b *Bsc) getGasPrice() (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	return b.ethClient.SuggestGasPrice(ctx)

}

func (b *Bsc) setPacketOpts() error {
	var curGasPrice *big.Int
	for {
		gasPrice, err := b.getGasPrice()
		if err != nil {
			return err
		}
		cmpRes := b.maxGasPrice.Cmp(gasPrice)
		if cmpRes == -1 {
			time.Sleep(TryGetGasPriceTimeInterval)
			continue
		} else {
			gasPriceUint := gasPrice.Int64()
			gasPriceUint += int64(float64(gasPriceUint) * b.tipCoefficient)
			curGasPrice = new(big.Int).SetInt64(gasPriceUint)
			break
		}
	}

	b.bindOpts.packetTransactOpts.GasPrice = curGasPrice
	return nil
}

func (b *Bsc) setClientOpts() error {
	var curGasPrice *big.Int
	for {
		gasPrice, err := b.getGasPrice()
		if err != nil {
			return err
		}
		cmpRes := b.maxGasPrice.Cmp(gasPrice)
		if cmpRes == -1 {
			continue
		} else {
			gasPriceUint := gasPrice.Int64()
			gasPriceUint += int64(float64(gasPriceUint) * b.tipCoefficient)
			curGasPrice = new(big.Int).SetInt64(gasPriceUint)
			break
		}
	}

	b.bindOpts.client.GasPrice = curGasPrice
	return nil
}

// ==================================================================================================================
// contract bind opts
type bindOpts struct {
	client             *bind.TransactOpts
	packetTransactOpts *bind.TransactOpts
}

func newBindOpts(cfg *ContractBindOptsCfg) (*bindOpts, error) {
	cliPriv, err := ethcrypto.HexToECDSA(cfg.ClientPrivKey)
	if err != nil {
		return nil, err
	}
	clientOpts, err := bind.NewKeyedTransactorWithChainID(cliPriv, new(big.Int).SetUint64(cfg.ChainID))
	if err != nil {
		return nil, err
	}
	clientOpts.GasLimit = cfg.GasLimit

	// packet transfer opts
	packPriv, err := ethcrypto.HexToECDSA(cfg.PacketPrivKey)
	if err != nil {
		return nil, err
	}
	packOpts, err := bind.NewKeyedTransactorWithChainID(packPriv, new(big.Int).SetUint64(cfg.ChainID))
	if err != nil {
		return nil, err
	}
	packOpts.GasLimit = cfg.GasLimit

	return &bindOpts{
		client:             clientOpts,
		packetTransactOpts: packOpts,
	}, nil
}

// ==================================================================================================================
// contract client group
type contractGroup struct {
	Packet   *contracts.Contract
	Client   *contracts.Contracts
	Transfer *transfer.Contracts
}

func newContractGroup(ethClient *ethclient.Client, cfgGroup *ContractCfgGroup) (*contractGroup, error) {
	packAddr := common.HexToAddress(cfgGroup.Packet.Addr)
	packetFilter, err := contracts.NewContract(packAddr, ethClient)
	if err != nil {
		return nil, err
	}

	clientAddr := common.HexToAddress(cfgGroup.Client.Addr)
	clientFilter, err := contracts.NewContracts(clientAddr, ethClient)
	if err != nil {
		return nil, err
	}
	transferAddress := common.HexToAddress(cfgGroup.Transfer.Addr)
	transferFilter, err := transfer.NewContracts(transferAddress, ethClient)
	if err != nil {
		return nil, err
	}

	return &contractGroup{
		Packet:   packetFilter,
		Client:   clientFilter,
		Transfer: transferFilter,
	}, nil
}
