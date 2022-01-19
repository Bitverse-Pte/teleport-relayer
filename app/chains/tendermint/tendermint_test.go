package tendermint

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"

	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/teleport-network/teleport-sdk-go/client"
	transfertypes "github.com/teleport-network/teleport/x/xibc/apps/transfer/types"
	xibctendermint "github.com/teleport-network/teleport/x/xibc/clients/light-clients/tendermint/types"
	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"

	"github.com/teleport-network/teleport-relayer/app/types"
)

// editable settings for test
const (
	GrpcUrl  = "10.41.20.10:9090"
	GrpcUrl2 = "127.0.0.1:19090"
	ChainId  = "teleport_9000-1"
	ChainId2 = "teleport_8544154630257-1"
)

var (
	testAcc1 = testAcc{
		name:     "node0",
		addr:     "teleport1pltgk26la3997f0rfaqcn7hxxpdqc836wda63x",
		mnemonic: "install rebel left tree aim capital truth rival demise auto enlist vote hybrid spare trick bounce cave forum amount track audit cake burst quick",
	}

	testAcc2 = testAcc{
		name:     "node0",
		addr:     "teleport1c892h6z5yslz4tj75cp63lu4cdawm378nqedy5",
		mnemonic: "turkey zebra curve enlist spring element region utility surge clip spray twist goddess decade october welcome beyond almost this february fiber chief cradle prison",
	}
)

type testAcc struct {
	name     string
	addr     string
	mnemonic string
}

func newClient() (*client.TeleportClient, error) {
	c, err := client.NewClient(GrpcUrl, ChainId)
	if err != nil {
		return nil, err
	}
	err = c.WithKeyring(keyring.NewInMemory(c.GetCtx().KeyringOptions...)).
		ImportMnemonic(testAcc1.name, testAcc1.mnemonic)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newClient2() (*client.TeleportClient, error) {
	c, err := client.NewClient(GrpcUrl2, ChainId2)
	if err != nil {
		return nil, err
	}
	if err = c.WithKeyring(keyring.NewInMemory(c.GetCtx().KeyringOptions...)).ImportMnemonic(testAcc2.name, testAcc2.mnemonic); err != nil {
		return nil, err
	}
	return c, nil
}

func Test1GetValSetByHeight(t *testing.T) {
	clientState := xibctendermint.ClientState{}
	cliStateBytes, _ := ioutil.ReadFile("")
	_ = json.Unmarshal(cliStateBytes, &clientState)
	protoByte, _ := proto.Marshal(&clientState)
	fmt.Println(string(protoByte))
	cli, _ := newClient()

	//req := new(tmservice.GetLatestBlockRequest)
	req := tmservice.GetBlockByHeightRequest{
		Height: 177411,
	}

	block, err := cli.TMServiceQuery.GetBlockByHeight(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	var height = block.Block.Header.Height
	txs := block.Block.GetData().Txs

	tx, _ := cli.TxClient.GetTx(context.Background(), &tx.GetTxRequest{
		Hash: hex.EncodeToString(tmhash.Sum(txs[0])),
	})
	events := tx.TxResponse.Logs[0].Events

	datas := getEventsValues(types.EventTypeSendPacket, "data", events)

	fmt.Println(txs)
	fmt.Println("height===", height)

	for _, data := range datas {
		fungibleTokenPacketData := &transfertypes.FungibleTokenPacketData{}
		d := data[1 : len(data)-1]
		var buf bytes.Buffer
		for _, v := range d {
			buf.WriteString(fmt.Sprintf("%q", v))
		}
		if err := fungibleTokenPacketData.Unmarshal([]byte(d)); err != nil {
			if fungibleTokenPacketData == nil {
				continue
			}

		}

		// fmt.Println("=============",fungibleTokenPacketData)
	}

	// ptt := packettypes.Packet{
	// 	Sequence:1,
	// 	// identifies the port on the sending chain and destination chain.
	// 	Port:"0x782ce34F9e84286669753bca9c9042DF19Dc60f7",
	// 	// identifies the chain id of the sending chain.
	// 	SourceChain:"testCreateClientB",
	// 	// identifies the chain id of the receiving chain.
	// 	DestinationChain:"testCreateClientA",
	// 	// identifies the chain id of the relay chain.
	// 	RelayChain:"",
	// 	// actual opaque bytes transferred directly to the application module
	// 	Data:[]byte("hello"),
	// }
	// pt := packettypes.MsgRecvPacket{
	// 	Packet: ptt,
	// 	ProofCommitment:[]byte("xxxx"),
	// 	ProofHeight:clienttypes.Height{
	// 		RevisionNumber:1,
	// 		RevisionHeight:1,
	// 	},
	// 	Signer:"teleport1pltgk26la3997f0rfaqcn7hxxpdqc836wda63x",
	// 	//Signer:"teleport1c892h6z5yslz4tj75cp63lu4cdawm378nqedy5",
	// }
	// ptres,err := cli.RecvPacket(pt)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("res===",ptres)
	// }
}
func getEventsValues(typ, key string, stringEvents sdk.StringEvents) []string {
	var res []string
	for _, e := range stringEvents {
		if e.Type == typ {
			for _, attr := range e.Attributes {
				if attr.Key == key {
					v := attr.Value
					fmt.Println(v)
					res = append(res, attr.Value)
				}
			}
		}
	}
	fmt.Println(res)
	return res
}

func Test2GetValSetByHeight(t *testing.T) {
	cli, err := newClient2()
	assert.NoError(t, err)

	res, err := cli.TMServiceQuery.GetValidatorSetByHeight(
		context.Background(),
		&tmservice.GetValidatorSetByHeightRequest{
			Height: 3,
			Pagination: &query.PageRequest{
				Offset: uint64(2),
				Limit:  5,
			}, //TODO
		},
	)
	assert.NoError(t, err)
	fmt.Println("res", res.Validators[0].Address)

	stateRes, err := cli.XIBCClientQuery.ClientState(
		context.Background(),
		&clienttypes.QueryClientStateRequest{ChainName: "testCreateClientA"},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("pf======", string(stateRes.Proof))
	fmt.Println("ph======", stateRes.ProofHeight)

	// ptt := packettypes.Packet{
	// 	Sequence:1,
	// 	// identifies the port on the sending chain and destination chain.
	// 	Port:"0x782ce34F9e84286669753bca9c9042DF19Dc60f7",
	// 	// identifies the chain id of the sending chain.
	// 	SourceChain:"testCreateClientB",
	// 	// identifies the chain id of the receiving chain.
	// 	DestinationChain:"testCreateClientA",
	// 	// identifies the chain id of the relay chain.
	// 	RelayChain:"",
	// 	// actual opaque bytes transferred directly to the application module
	// 	Data:[]byte("hello"),
	// }
	// pt := packettypes.MsgRecvPacket{
	// 	Packet: ptt,
	// 	ProofCommitment:[]byte("xxxx"),
	// 	ProofHeight:clienttypes.Height{
	// 		RevisionNumber:1,
	// 		RevisionHeight:1,
	// 	},
	// 	//Signer:"teleport1pltgk26la3997f0rfaqcn7hxxpdqc836wda63x",
	// 	Signer:"teleport1c892h6z5yslz4tj75cp63lu4cdawm378nqedy5",
	// }
	// ptres,err := cli.RecvPacket(pt)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("res===",ptres)
	// }
}
