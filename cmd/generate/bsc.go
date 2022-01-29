package generate

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"

	"golang.org/x/crypto/sha3"

	"github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"

	xibcbsc "github.com/teleport-network/teleport/x/xibc/clients/light-clients/bsc/types"
	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"

	"github.com/teleport-network/teleport-relayer/app/config"
)

const epoch = uint64(200)

func generateBscJson(bsc *config.Bsc, codec *codec.ProtoCodec, logger *logrus.Entry) {
	fmt.Printf("%+v \n", bsc)
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	rpcClient, err := rpc.DialContext(ctx, bsc.URI)
	if err != nil {
		logger.Fatal(err)
	}
	ethClient := ethclient.NewClient(rpcClient)
	latestHeight, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		logger.Fatal(err)
	}

	startHeight := latestHeight - latestHeight%epoch
	logger.Info("bsc height = ", startHeight)

	blockRes, err := ethClient.BlockByNumber(
		context.Background(),
		new(big.Int).SetUint64(startHeight),
	)
	if err != nil {
		logger.Fatal(err)
	}

	blockHeader := blockRes.Header()
	genesisValidators, err := xibcbsc.ParseValidators(blockHeader.Extra)
	if err != nil {
		logger.Fatal(err)
	}
	header := &xibcbsc.BscHeader{
		ParentHash:  blockHeader.ParentHash,
		UncleHash:   blockHeader.UncleHash,
		Coinbase:    blockHeader.Coinbase,
		Root:        blockHeader.Root,
		TxHash:      blockHeader.TxHash,
		ReceiptHash: blockHeader.ReceiptHash,
		Bloom:       xibcbsc.BytesToBloom(blockHeader.Bloom.Bytes()),
		Difficulty:  blockHeader.Difficulty,
		Number:      blockHeader.Number,
		GasLimit:    blockHeader.GasLimit,
		GasUsed:     blockHeader.GasUsed,
		Time:        blockHeader.Time,
		Extra:       blockHeader.Extra,
		MixDigest:   blockHeader.MixDigest,
		Nonce:       xibcbsc.EncodeNonce(blockHeader.Nonce.Uint64()),
	}
	number := clienttypes.NewHeight(0, header.Number.Uint64())
	hash := common.FromHex(bsc.Contracts.Packet.Addr)
	fmt.Println("cfg.Bsc.Contracts.Packet.Addr=", bsc.Contracts.Packet.Addr)
	clientState := &xibcbsc.ClientState{
		Header:          header.ToHeader(),
		ChainId:         bsc.ChainID,
		Epoch:           epoch,
		BlockInteval:    3,
		Validators:      genesisValidators,
		ContractAddress: hash,
		TrustingPeriod:  60 * 60 * 24 * 100,
	}
	consensusState := &xibcbsc.ConsensusState{
		Timestamp: header.Time,
		Height:    number,
		Root:      header.Root[:],
	}
	signer := ecrecover(header.ToHeader(), big.NewInt(int64(bsc.ChainID)))
	equal := bytes.Equal(header.Coinbase.Bytes(), signer.Bytes())
	if !equal {
		logger.Fatal("header.Coinbase")
	}

	clientStateBytes, err := codec.MarshalInterfaceJSON(clientState)
	if err != nil {
		logger.Fatal(err)
	}
	clientStateFilename := fmt.Sprintf("%s_clientState.json", bsc.ChainName)
	WriteCreateClientFiles(clientStateFilename, string(clientStateBytes))

	consensusStateBytes, err := codec.MarshalInterfaceJSON(consensusState)
	if err != nil {
		logger.Fatal(err)
	}
	consensusStateFilename := fmt.Sprintf("%s_consensusState.json", bsc.ChainName)
	WriteCreateClientFiles(consensusStateFilename, string(consensusStateBytes))
}

const extraSeal = 65

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header xibcbsc.Header, chainId *big.Int) common.Address {
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < extraSeal {
		return common.Address{}
	}
	signature := header.Extra[len(header.Extra)-extraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(sealHash(header, chainId).Bytes(), signature)
	if err != nil {
		return common.Address{}
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	return signer
}

// sealHash returns the hash of a block prior to it being sealed.
func sealHash(header xibcbsc.Header, chainId *big.Int) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header, chainId)
	hasher.Sum(hash[:0])
	return hash
}
func encodeSigHeader(w io.Writer, header xibcbsc.Header, chainId *big.Int) {
	err := rlp.Encode(w, []interface{}{
		chainId,
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Height.RevisionHeight,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-65], // this will panic if extra is too short, should check before calling encodeSigHeader
		header.MixDigest,
		header.Nonce,
	})
	if err != nil {
		panic("can't encode: " + err.Error())
	}
}
