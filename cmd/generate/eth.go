package generate

import (
	"context"
	"fmt"
	"math/big"

	"github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	xibceth "github.com/teleport-network/teleport/x/xibc/clients/light-clients/eth/types"
	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/services/tendermint"
)

func generateETHJson(cfg *config.ChainCfg, tmClient *tendermint.Tendermint, logger *logrus.Entry) {
	fmt.Printf("%+v", cfg.Eth)
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	rpcClient, err := rpc.DialContext(ctx, cfg.Eth.URI)
	if err != nil {
		logger.Fatal(err)
	}
	ethClient := ethclient.NewClient(rpcClient)
	latestHeight, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		logger.Fatal(err)
	}
	startHeight := latestHeight - 10
	logger.Info("eth height = ", startHeight)

	blockRes, err := ethClient.BlockByNumber(
		context.Background(),
		new(big.Int).SetUint64(startHeight),
	)
	if err != nil {
		logger.Fatal(err)
	}

	blockHeader := blockRes.Header()

	header := &xibceth.EthHeader{
		ParentHash:  blockHeader.ParentHash,
		UncleHash:   blockHeader.UncleHash,
		Coinbase:    blockHeader.Coinbase,
		Root:        blockHeader.Root,
		TxHash:      blockHeader.TxHash,
		ReceiptHash: blockHeader.ReceiptHash,
		Bloom:       blockHeader.Bloom,
		Difficulty:  blockHeader.Difficulty,
		Number:      blockHeader.Number,
		GasLimit:    blockHeader.GasLimit,
		GasUsed:     blockHeader.GasUsed,
		Time:        blockHeader.Time,
		Extra:       blockHeader.Extra,
		MixDigest:   blockHeader.MixDigest,
		Nonce:       blockHeader.Nonce,
		BaseFee:     blockHeader.BaseFee,
	}
	number := clienttypes.NewHeight(0, header.Number.Uint64())
	hash := common.FromHex(cfg.Eth.Contracts.Packet.Addr)
	fmt.Println("cfg.Eth.Contracts.Packet.Addr=", cfg.Eth.Contracts.Packet.Addr)
	clientState := &xibceth.ClientState{
		Header:          header.ToHeader(),
		ChainId:         cfg.Eth.ChainID,
		ContractAddress: hash,
		TrustingPeriod:  60 * 60 * 24 * 100,
		TimeDelay:       0,
		BlockDelay:      7,
	}
	consensusState := &xibceth.ConsensusState{
		Timestamp: header.Time,
		Number:    number,
		Root:      header.Root[:],
	}

	clientStateBytes, err := tmClient.Codec.MarshalJSON(clientState)
	if err != nil {
		logger.Fatal(err)
	}

	clientStateStr := string(clientStateBytes)
	clientStateStr = EthClientStatePrefix + clientStateStr[1:]
	clientStateFilename := fmt.Sprintf("%s_clientState.json", cfg.Eth.ChainName)
	WriteCreateClientFiles(clientStateFilename, clientStateStr)

	consensusStateBytes, err := tmClient.Codec.MarshalJSON(consensusState)
	if err != nil {
		logger.Fatal(err)
	}

	consensusStateStr := string(consensusStateBytes)
	consensusStateStr = EthConsensusStatePrefix + consensusStateStr[1:]
	consensusStateFilename1 := fmt.Sprintf("%s_consensusState.json", cfg.Eth.ChainName)
	WriteCreateClientFiles(consensusStateFilename1, consensusStateStr)
}
