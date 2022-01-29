package generate

import (
	"context"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	xibceth "github.com/teleport-network/teleport/x/xibc/clients/light-clients/eth/types"
	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"

	"github.com/teleport-network/teleport-relayer/app/config"
)

func generateETHJson(eth *config.Eth, codec *codec.ProtoCodec, logger *logrus.Entry) {
	fmt.Printf("%+v \n", eth)
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	rpcClient, err := rpc.DialContext(ctx, eth.URI)
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
	hash := common.FromHex(eth.Contracts.Packet.Addr)
	fmt.Println("cfg.Eth.Contracts.Packet.Addr=", eth.Contracts.Packet.Addr)
	clientState := &xibceth.ClientState{
		Header:          header.ToHeader(),
		ChainId:         eth.ChainID,
		ContractAddress: hash,
		TrustingPeriod:  60 * 60 * 24 * 100,
		TimeDelay:       0,
		BlockDelay:      7,
	}
	consensusState := &xibceth.ConsensusState{
		Timestamp: header.Time,
		Height:    number,
		Root:      header.Root[:],
	}

	clientStateBytes, err := codec.MarshalInterfaceJSON(clientState)
	if err != nil {
		logger.Fatal(err)
	}
	clientStateFilename := fmt.Sprintf("%s_clientState.json", eth.ChainName)
	WriteCreateClientFiles(clientStateFilename, string(clientStateBytes))

	consensusStateBytes, err := codec.MarshalInterfaceJSON(consensusState)
	if err != nil {
		logger.Fatal(err)
	}
	consensusStateFilename := fmt.Sprintf("%s_consensusState.json", eth.ChainName)
	WriteCreateClientFiles(consensusStateFilename, string(consensusStateBytes))
}
