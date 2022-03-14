package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/teleport-network/teleport-relayer/app/chains/eth/contracts"
	"github.com/teleport-network/teleport-relayer/app/chains/eth/contracts/transfer"
	"github.com/teleport-network/teleport-relayer/app/interfaces"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"

	xibceth "github.com/teleport-network/teleport/x/xibc/clients/light-clients/eth/types"
	xibctendermint "github.com/teleport-network/teleport/x/xibc/clients/light-clients/tendermint/types"
	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"
	"github.com/teleport-network/teleport/x/xibc/core/host"
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"
	"github.com/teleport-network/teleport/x/xibc/exported"

	"github.com/teleport-network/teleport-relayer/app/types"
	"github.com/teleport-network/teleport-relayer/app/types/errors"
)

var _ interfaces.IChain = new(Eth)

const CtxTimeout = 100 * time.Second
const TryGetGasPriceTimeInterval = 10 * time.Second
const RetryTimeout = 15 * time.Second
const RetryTimes = 20

var (
	Uint64, _  = abi.NewType("uint64", "", nil)
	Bytes32, _ = abi.NewType("bytes32", "", nil)
	Bytes, _   = abi.NewType("bytes", "", nil)
	String, _  = abi.NewType("string", "", nil)
)

type Eth struct {
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
}

func NewEth(config *ChainConfig) (interfaces.IChain, error) {
	return newEth(config)
}

func newEth(config *ChainConfig) (*Eth, error) {
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

	return &Eth{
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
	}, nil
}

func (eth *Eth) ClientUpdateValidate(revisionHeight, delayHeight, updateHeight uint64) (uint64, error) {

	return updateHeight, nil
}

func (eth *Eth) TransferERC20(transferData transfer.TransferDataTypesERC20TransferData) error {
	resultTx := &types.ResultTx{}
	if err := eth.setPacketOpts(); err != nil {
		return err
	}
	result, err := eth.contracts.Transfer.SendTransferERC20(eth.bindOpts.packetTransactOpts, transferData)
	if err != nil {
		return err
	}
	resultTx.GasUsed += int64(result.Gas())
	resultTx.Hash = resultTx.Hash + "," + result.Hash().String()
	return eth.reTryEthResult(resultTx.Hash, 0)
}

func (eth *Eth) RelayPackets(msgs []sdk.Msg) (string, error) {
	resultTx := &types.ResultTx{}
	var packetDetail []string
	for _, d := range msgs {
		switch msg := d.(type) {
		case *packettypes.MsgRecvPacket:
			packetMsg := fmt.Sprintf("srcChain:%v,dest%v,sequence:%v chainType:packet", msg.Packet.SourceChain, msg.Packet.DestinationChain, msg.Packet.Sequence)
			packetDetail = append(packetDetail, packetMsg)
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
			if err := eth.setPacketOpts(); err != nil {
				return packetMsg, err
			}
			result, err := eth.contracts.Packet.RecvPacket(
				eth.bindOpts.packetTransactOpts,
				tmpPack,
				msg.ProofCommitment,
				height,
			)
			if err != nil {
				return fmt.Sprintf("relayer tx hash:%v\n packet detail:%v", result.Hash().String(), packetMsg), err
			}
			resultTx.Hash += "," + result.Hash().String()
		case *packettypes.MsgAcknowledgement:
			packetMsg := fmt.Sprintf("srcChain:%v,dest%v,sequence:%v,chainType: Ack", msg.Packet.SourceChain, msg.Packet.DestinationChain, msg.Packet.Sequence)
			packetDetail = append(packetDetail, packetMsg)
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

			if err := eth.setPacketOpts(); err != nil {
				return packetMsg, err
			}

			result, err := eth.contracts.Packet.AcknowledgePacket(
				eth.bindOpts.packetTransactOpts,
				tmpPack, msg.Acknowledgement, msg.ProofAcked,
				height,
			)
			if err != nil {
				return packetMsg, err
			}
			resultTx.Hash += "," + result.Hash().String()
		}
	}
	resultTx.Hash = strings.Trim(resultTx.Hash, ",")
	if err := eth.reTryEthResult(resultTx.Hash, 0); err != nil {
		return fmt.Sprintf("relayer tx hash :%v\n,packet detail:%v", resultTx.Hash, strings.Join(packetDetail, ",")), err
	}
	return fmt.Sprintf("relayer tx hash :%v\n,packet detail:%v", resultTx.Hash, strings.Join(packetDetail, ",")), nil
}

func (eth *Eth) UpdateClient(header exported.Header, chainName string) error {
	h, ok := header.(*xibctendermint.Header)
	if !ok {
		return fmt.Errorf("invalid header type")
	}
	headerBytes, _ := h.Marshal()
	if err := eth.setClientOpts(); err != nil {
		return err
	}
	result, err := eth.contracts.Client.UpdateClient(eth.bindOpts.client, chainName, headerBytes)
	if err != nil {
		return err
	}
	if err := eth.reTryEthResult(result.Hash().String(), 0); err != nil {
		return err // TODO: warp
	}
	return nil
}

func (eth *Eth) BatchUpdateClient(headers []exported.Header, chainName string) error {
	return nil
}

func (eth *Eth) reTryEthResult(hash string, n uint64) error {
	if n == RetryTimes {
		return fmt.Errorf("retry %d times and return error", RetryTimes)
	}
	txStatus, err := eth.GetResult(hash)
	if err != nil {
		time.Sleep(RetryTimeout)
		return eth.reTryEthResult(hash, n+1)
	}
	if txStatus == 0 {
		return fmt.Errorf("txStatus == 0, tx failed")
	}
	return nil
}

func (eth *Eth) GetPackets(fromBlock, toBlock uint64, destChainType string) (*types.Packets, error) {
	bizPackets, err := eth.getPackets(fromBlock, toBlock)
	if err != nil {
		return nil, err
	}
	ackPackets, err := eth.getAckPackets(fromBlock, toBlock)
	if err != nil {
		return nil, err
	}
	packets := &types.Packets{
		BizPackets: bizPackets,
		AckPackets: ackPackets,
	}
	return packets, nil
}

func (eth *Eth) GetProof(sourChainName, destChainName string, sequence uint64, height uint64, typ string) ([]byte, error) {
	pkConstr := xibceth.NewProofKeyConstructor(sourChainName, destChainName, sequence)
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
	address := common.HexToAddress(eth.contractCfgGroup.Packet.Addr)
	result, err := eth.getProof(ctx, address, []string{hexutil.Encode(key)}, new(big.Int).SetUint64(height))
	if err != nil {
		return nil, err
	}

	var storageProof []*xibceth.StorageResult
	for _, sp := range result.StorageProof {
		tmpStorageProof := &xibceth.StorageResult{
			Key:   sp.Key,
			Value: hexutil.EncodeBig(sp.Value),
			Proof: sp.Proof,
		}
		storageProof = append(storageProof, tmpStorageProof)
	}
	nonce := hexutil.EncodeUint64(result.Nonce)
	balance := hexutil.EncodeBig(result.Balance)
	proof := &xibceth.Proof{
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

func (eth *Eth) GetCommitmentsPacket(sourChainName, destChainName string, sequence uint64) error {
	hashBytes, err := eth.contracts.Packet.Commitments(nil, host.PacketCommitmentKey(sourChainName, destChainName, sequence))
	if err != nil {
		return err
	}
	expectByte := make([]byte, 32)
	if bytes.Equal(expectByte, hashBytes[:]) {
		return fmt.Errorf("commitment does not exist")
	}
	return nil
}

func (eth *Eth) GetReceiptPacket(sourChainName, destChainName string, sequence uint64) (bool, error) {
	return eth.contracts.Packet.Receipts(nil, host.PacketReceiptKey(sourChainName, destChainName, sequence))
}

func (eth *Eth) GetBlockTimestamp(height uint64) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	blockRes, err := eth.ethClient.BlockByNumber(ctx, new(big.Int).SetUint64(height))
	if err != nil {
		return 0, err
	}
	return blockRes.Time(), nil
}

func (eth *Eth) GetBlockHeader(req *types.GetBlockHeaderReq) (exported.Header, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	blockRes, err := eth.ethClient.BlockByNumber(ctx, new(big.Int).SetUint64(req.LatestHeight))
	if err != nil {
		return nil, err
	}
	return &xibceth.Header{
		ParentHash:  blockRes.ParentHash().Bytes(),
		UncleHash:   blockRes.UncleHash().Bytes(),
		Coinbase:    blockRes.Coinbase().Bytes(),
		Root:        blockRes.Root().Bytes(),
		TxHash:      blockRes.TxHash().Bytes(),
		ReceiptHash: blockRes.ReceiptHash().Bytes(),
		Bloom:       blockRes.Bloom().Bytes(),
		Difficulty:  blockRes.Difficulty().Bytes(),
		Height: clienttypes.Height{
			RevisionNumber: req.RevisionNumber,
			RevisionHeight: req.LatestHeight,
		},
		GasLimit:  blockRes.GasLimit(),
		GasUsed:   blockRes.GasUsed(),
		Time:      blockRes.Time(),
		Extra:     blockRes.Extra(),
		MixDigest: blockRes.MixDigest().Bytes(),
		Nonce:     blockRes.Nonce(),
		BaseFee:   blockRes.BaseFee().Bytes(),
	}, nil

}

func (eth *Eth) GetLightClientState(chainName string) (exported.ClientState, error) {
	latestHeight, err := eth.contracts.Client.GetLatestHeight(nil, chainName)
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

func (eth *Eth) GetLightClientConsensusState(string, uint64) (exported.ConsensusState, error) {
	return nil, nil
}

func (eth *Eth) GetLatestHeight() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	return eth.ethClient.BlockNumber(ctx)
}

func (eth *Eth) GetLightClientDelayHeight(chainName string) (uint64, error) {
	return 0, nil
}

func (eth *Eth) GetLightClientDelayTime(chainName string) (uint64, error) {
	return 0, nil
}

func (eth *Eth) GetResult(hash string) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()

	cmnHash := common.HexToHash(hash)
	result, err := eth.ethClient.TransactionReceipt(ctx, cmnHash)
	if err != nil {
		return 0, err
	}
	return result.Status, nil
}

func (eth *Eth) ChainName() string {
	return eth.chainName
}

func (eth *Eth) UpdateClientFrequency() uint64 {
	return eth.updateClientFrequency
}

func (eth *Eth) ChainType() string {
	return eth.chainType
}

func (eth *Eth) getProof(ctx context.Context, account common.Address, keys []string, blockNumber *big.Int) (*gethclient.AccountResult, error) {
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
	if err := eth.gethRpcCli.CallContext(ctx, &res, "eth_getProof", account, keys, toBlockNumArg(blockNumber)); err != nil {
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
func (eth *Eth) getPackets(fromBlock, toBlock uint64) ([]packettypes.Packet, error) {
	if strings.Contains(eth.queryFilter, types.Packet) {
		return nil, nil
	}
	address := common.HexToAddress(eth.contractCfgGroup.Packet.Addr)
	topic := eth.contractCfgGroup.Packet.Topic
	logs, err := eth.getLogs(address, topic, fromBlock, toBlock)
	if err != nil {
		return nil, err
	}
	var bizPackets []packettypes.Packet
	for _, log := range logs {
		packSent, err := eth.contracts.Packet.ParsePacketSent(log)
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
func (eth *Eth) getAckPackets(fromBlock, toBlock uint64) ([]types.AckPacket, error) {
	if strings.Contains(eth.queryFilter, types.Ack) {
		return nil, nil
	}
	address := common.HexToAddress(eth.contractCfgGroup.AckPacket.Addr)
	topic := eth.contractCfgGroup.AckPacket.Topic
	logs, err := eth.getLogs(address, topic, fromBlock, toBlock)
	if err != nil {
		return nil, err
	}

	var ackPackets []types.AckPacket
	for _, log := range logs {
		ackWritten, err := eth.contracts.Packet.ParseAckWritten(log)
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

func (eth *Eth) getLogs(address common.Address, topic string, fromBlock, toBlock uint64) ([]ethtypes.Log, error) {
	filter := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []common.Address{address},
		Topics:    [][]common.Hash{{ethcrypto.Keccak256Hash([]byte(topic))}},
	}
	return eth.ethClient.FilterLogs(context.Background(), filter)
}

func (eth *Eth) getGasPrice() (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	return eth.ethClient.SuggestGasPrice(ctx)

}

func (eth *Eth) setPacketOpts() error {
	var curGasPrice *big.Int
	for {
		gasPrice, err := eth.getGasPrice()
		if err != nil {
			return err
		}
		cmpRes := eth.maxGasPrice.Cmp(gasPrice)
		if cmpRes == -1 {
			time.Sleep(TryGetGasPriceTimeInterval)
			continue
		} else {
			gasPriceUint := gasPrice.Int64()
			gasPriceUint += int64(float64(gasPriceUint) * eth.tipCoefficient)
			curGasPrice = new(big.Int).SetInt64(gasPriceUint)
			break
		}
	}

	eth.bindOpts.packetTransactOpts.GasPrice = curGasPrice
	return nil
}

func (eth *Eth) setClientOpts() error {
	var curGasPrice *big.Int
	for {
		gasPrice, err := eth.getGasPrice()
		if err != nil {
			return err
		}
		cmpRes := eth.maxGasPrice.Cmp(gasPrice)
		if cmpRes == -1 {
			continue
		} else {
			gasPriceUint := gasPrice.Int64()
			gasPriceUint += int64(float64(gasPriceUint) * eth.tipCoefficient)
			curGasPrice = new(big.Int).SetInt64(gasPriceUint)
			break
		}
	}

	eth.bindOpts.client.GasPrice = curGasPrice
	return nil
}

// func (eth *Eth) setTransferOpts() error {
// 	var curGasPrice *big.Int
// 	for {
// 		gasPrice, err := eth.getGasPrice()
// 		if err != nil {
// 			return err
// 		}
// 		if eth.maxGasPrice.Cmp(gasPrice) != -1 {
// 			break
// 		}
// 	}

// 	eth.bindOpts.client.GasPrice = curGasPrice
// 	return nil
// }

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
