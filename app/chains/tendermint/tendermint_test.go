package tendermint

import (
	"context"
	"encoding/hex"
	"fmt"
	"sync"
	"testing"

	"github.com/teleport-network/teleport-relayer/app/utils"

	"github.com/cosmos/cosmos-sdk/types/query"

	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/teleport-network/teleport-sdk-go/client"
	xibctendermint "github.com/teleport-network/teleport/x/xibc/clients/light-clients/tendermint/types"

	"github.com/teleport-network/teleport-relayer/app/types"
)

// editable settings for test
const (
	GrpcUrl = "abd46ec6e28754f0ab2aae29deaa0c11-1510914274.ap-southeast-1.elb.amazonaws.com:9090"
	ChainId = "teleport_7001-1"

	localGrpc    = "localhost:9090"
	localChainId = "teleport_9000-1"
)

type testAcc struct {
	name     string
	addr     string
	mnemonic string
}

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

func TestGetHashByEvent(t *testing.T) {
	cli, err := newClient()
	require.NoError(t, err)

	Pagination := &query.PageRequest{
		Key:        []byte(""),
		Limit:      200,
		Offset:     0,
		CountTotal: false,
		Reverse:    false,
	}
	res, err := cli.TxClient.GetTxsEvent(context.Background(), &tx.GetTxsEventRequest{Events: []string{"xibc.core.packet.v1.EventRecvPacket.src_chain='rinkeby'"}, Pagination: Pagination, OrderBy: 1})
	require.NoError(t, err)
	t.Log(len(res.Txs))
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

func TestGenTendermintHeader(t *testing.T) {
	c, err := client.NewClient(GrpcUrl, ChainId)
	require.NoError(t, err)
	err = c.WithKeyring(keyring.NewInMemory(c.GetCtx().KeyringOptions...)).ImportMnemonic(testAcc1.name, testAcc1.mnemonic)
	require.NoError(t, err)
	tendermint := Tendermint{
		TeleportSDK: c,
		Codec:       MakeCodec(),
	}
	req := &types.GetBlockHeaderReq{
		LatestHeight:   4926,
		TrustedHeight:  4882,
		RevisionNumber: 1,
	}
	header, err := tendermint.GetBlockHeader(req)
	require.NoError(t, err)
	h, ok := header.(*xibctendermint.Header)
	require.True(t, ok)
	headerBytes, _ := h.Marshal()
	hexHeader := hex.EncodeToString(headerBytes)
	fmt.Println(hexHeader)
}

func newTendermintClient(grpc, id string) *Tendermint {
	cdc := MakeCodec()
	cli, err := client.NewClient(grpc, id)
	if err != nil {
		panic(fmt.Errorf("tendermint new client error:%+v", err))
	}
	if err := cli.WithKeyring(keyring.NewInMemory(cli.GetCtx().KeyringOptions...)).ImportMnemonic(testAcc1.name, testAcc1.mnemonic); err != nil {
		panic(fmt.Errorf("tendermint cli.WithKeyring error:%+v", err))
	}
	address, err := cli.Key(testAcc1.name)
	if err != nil {
		panic(fmt.Errorf("cli.Key error:%+v", err))
	}
	simulationClient, err := client.NewClient(grpc, id)
	if err != nil {
		panic(fmt.Errorf("tendermint new client error:%+v", err))
	}
	if err := simulationClient.WithKeyring(keyring.NewInMemory(simulationClient.GetCtx().KeyringOptions...)).ImportMnemonic(testAcc1.name, testAcc1.mnemonic); err != nil {
		panic(fmt.Errorf("tendermint cli.WithKeyring error:%+v", err))
	}
	simulationClient.WithAccountRetrieverCache(cli.GetAccountRetriever().Cache)

	return &Tendermint{
		chainType:             types.Tendermint,
		chainName:             "teleport",
		Codec:                 cdc,
		TeleportSDK:           cli,
		SimulationClient:      simulationClient,
		updateClientFrequency: 0,
		address:               address,
		l:                     new(sync.Mutex),
	}
}

func TestGetPacketsByHash(t *testing.T) {
	c := newTendermintClient(localGrpc, localChainId)
	packet, err := c.GetPacketsByHash("")
	require.NoError(t, err)
	require.NotNil(t, packet)
}

func TestGetPacketsByHeight(t *testing.T) {
	c := newTendermintClient(GrpcUrl, ChainId)
	packet, err := c.GetPackets(1249646, 1249646, "")
	require.NoError(t, err)
	require.NotNil(t, packet)

	for _, v := range packet.BizPackets {
		t.Log("srcChain:", v.GetSrcChain())
		t.Log("destChain:", v.GetDstChain())
		t.Log("sequence:", v.GetSequence())

		var transferData packettypes.TransferData
		err = transferData.ABIDecode(v.TransferData)
		require.NoError(t, err)
		require.NotNil(t, transferData)
		t.Log("TransferData: ", transferData.String())

		if len(v.CallData) != 0 {
			var callData packettypes.CallData
			err = callData.ABIDecode(v.CallData)
			require.NoError(t, err)
			require.NotNil(t, callData)
			t.Log("CallData: ", callData.String())
		}
	}

	check, err := utils.BridgeTimeLimitCheck(packet.BizPackets, "https://bridge.qa.davionlabs.com/bridge/status")
	require.NoError(t, err)

	t.Log(check)
}
