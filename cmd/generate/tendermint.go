package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/teleport-network/teleport-relayer/app/chains/tendermint"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sirupsen/logrus"

	xibceth "github.com/teleport-network/teleport/x/xibc/clients/light-clients/eth/types"
	xibctendermint "github.com/teleport-network/teleport/x/xibc/clients/light-clients/tendermint/types"
	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"
	commitmenttypes "github.com/teleport-network/teleport/x/xibc/core/commitment/types"
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"

	"github.com/tendermint/tendermint/libs/bytes"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
)

type TendermintConsensusState struct {
	Timestamp          Timestamp `json:"timestamp"`
	Root               string    `json:"root"`
	NextValidatorsHash string    `json:"nextValidatorsHash"`
}

type Timestamp struct {
	Secs  int64 `json:"secs"`
	Nanos int64 `json:"nanos"`
}

func makeCodec() *codec.ProtoCodec {
	ir := codectypes.NewInterfaceRegistry()
	clienttypes.RegisterInterfaces(ir)
	govtypes.RegisterInterfaces(ir)
	xibctendermint.RegisterInterfaces(ir)
	xibceth.RegisterInterfaces(ir)
	packettypes.RegisterInterfaces(ir)
	ir.RegisterInterface("cosmos.v1beta1.Msg", (*sdk.Msg)(nil))
	tx.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)
	return codec.NewProtoCodec(ir)
}

func generateTendermintHex(
	client *tendermint.Tendermint,
	height int64,
	chainName string,
	logger *logrus.Entry,
) {
	res, err := client.TeleportSDK.TMServiceQuery.GetBlockByHeight(
		context.Background(),
		&tmservice.GetBlockByHeightRequest{
			Height: height,
		},
	)
	if err != nil {
		return
	}
	tmHeader := res.Block.Header
	si := &tmtypes.SignedHeader{
		Header: &res.Block.Header,
	}
	prHeader := xibctendermint.Header{
		SignedHeader: si,
	}
	revisionNumber := int(prHeader.GetHeight().GetRevisionNumber())
	revisionHeight := prHeader.GetHeight().GetRevisionHeight()
	clientState := &xibctendermint.ClientState{
		ChainId: tmHeader.ChainID,
		TrustLevel: xibctendermint.Fraction{
			Numerator:   1,
			Denominator: 3,
		},
		TrustingPeriod:  100 * 24 * 60 * 60,
		UnbondingPeriod: 110 * 24 * 60 * 60,
		MaxClockDrift:   10,
		LatestHeight: clienttypes.Height{
			RevisionNumber: uint64(revisionNumber),
			RevisionHeight: revisionHeight,
		},
		MerklePrefix: commitmenttypes.MerklePrefix{
			KeyPrefix: []byte("xibc"),
		},
		TimeDelay: 0,
	}
	//consensusState := xibctendermint.ConsensusState{
	//	Timestamp:          tmHeader.Time,
	//	Root:               bytes.HexBytes(tmHeader.AppHash),
	//	NextValidatorsHash: bytes.HexBytes(tmHeader.NextValidatorsHash), // TODO check
	//}

	consensusState := TendermintConsensusState{
		Timestamp: Timestamp{
			Secs:  tmHeader.Time.Unix(),
			Nanos: 0,
		},
		Root:               bytes.HexBytes(tmHeader.AppHash).String(),
		NextValidatorsHash: bytes.HexBytes(tmHeader.NextValidatorsHash).String(),
	}
	clientStateBytes, err := json.Marshal(clientState)
	// clientStateBytes, err := proto.Marshal(clientState)
	if err != nil {
		logger.Fatal("marshal eth clientState error: ", err)
	}
	// write file
	clientStateFilename := fmt.Sprintf("%s_clientState.json", chainName)
	WriteCreateClientFiles(clientStateFilename, string(clientStateBytes))

	clientStateFilename2 := fmt.Sprintf("%s_clientState.txt", chainName)
	WriteCreateClientFiles(clientStateFilename2, hexutil.Encode(clientStateBytes)[2:])
	fmt.Println("clientState: ", hexutil.Encode(clientStateBytes)[2:])
	consensusStateBytes, err := json.Marshal(&consensusState)
	// consensusStateBytes, err := json.Marshal(consensusState)
	if err != nil {
		logger.Fatal(err)
	}
	consensusStateFilename := fmt.Sprintf("%s_consensusState.json", chainName)
	WriteCreateClientFiles(consensusStateFilename, string(consensusStateBytes))
	consensusStateFilename2 := fmt.Sprintf("%s_consensusState.txt", chainName)
	WriteCreateClientFiles(consensusStateFilename2, hexutil.Encode(consensusStateBytes)[2:])
	fmt.Println("consensusState: ", hexutil.Encode(consensusStateBytes)[2:])
}

func generateTendermintJson(
	client *tendermint.Tendermint,
	height int64,
	chainName string,
) {
	res, err := client.TeleportSDK.TMServiceQuery.GetBlockByHeight(
		context.Background(),
		&tmservice.GetBlockByHeightRequest{Height: height},
	)
	if err != nil {
		fmt.Println("QueryBlock fail:  ", err)
	}
	tmHeader := res.Block.Header
	si := &tmtypes.SignedHeader{ // TODO check
		Header: &res.Block.Header,
	}
	prHeader := xibctendermint.Header{
		SignedHeader: si,
	}

	lastHeight := clienttypes.NewHeight(
		prHeader.GetHeight().GetRevisionNumber(),
		prHeader.GetHeight().GetRevisionHeight(),
	)

	var clientState = &xibctendermint.ClientState{
		ChainId: tmHeader.ChainID,
		TrustLevel: xibctendermint.Fraction{
			Numerator:   1,
			Denominator: 3,
		},
		TrustingPeriod:  time.Hour * 24 * 100,
		UnbondingPeriod: time.Hour * 24 * 110,
		MaxClockDrift:   time.Second * 10,
		LatestHeight:    lastHeight,
		ProofSpecs:      commitmenttypes.GetSDKSpecs(),
		MerklePrefix:    commitmenttypes.MerklePrefix{KeyPrefix: []byte(xibcTendermintMerklePrefix)},
		TimeDelay:       0,
	}
	validatorSet, err := client.GetValidator(height)
	if err != nil {
		panic(err)
	}
	var consensusState = &xibctendermint.ConsensusState{
		Timestamp:          tmHeader.Time,
		Root:               []byte(xibcTendermintRoot),
		NextValidatorsHash: validatorSet.Hash(),
	}

	clientStateBytes, err := client.Codec.MarshalInterfaceJSON(clientState)
	if err != nil {
		panic(err)
	}
	// write file
	clientStateFilename := fmt.Sprintf("%s_clientState.json", chainName)
	WriteCreateClientFiles(clientStateFilename, string(clientStateBytes))
	consensusStateBytes, err := client.Codec.MarshalInterfaceJSON(consensusState)
	if err != nil {
		panic(err)
	}
	consensusStateFilename := fmt.Sprintf("%s_consensusState.json", chainName)
	WriteCreateClientFiles(consensusStateFilename, string(consensusStateBytes))
}
