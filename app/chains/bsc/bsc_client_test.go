package bsc

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"

	transfertypes "github.com/teleport-network/teleport/x/xibc/apps/transfer/types"
	"github.com/teleport-network/teleport/x/xibc/core/host"

	"github.com/teleport-network/teleport-relayer/app/types"
)

const (
	testUrl = "https://data-seed-prebsc-1-s1.binance.org:8545/"
	testId  = 97
)

func TestNewBsc(t *testing.T) {
	bscClient := newBscClient(t)

	latestHeight, err := bscClient.GetLatestHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(latestHeight)
}

func TestGetPackets(t *testing.T) {
	bscClient := newBscClient(t)

	packets, err := bscClient.GetPackets(18666603, 18666603, "")
	require.NoError(t, err)
	require.NotNil(t, packets.BizPackets)
	var data transfertypes.FungibleTokenPacketData
	err = data.DecodeBytes(packets.BizPackets[0].DataList[0])
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestGetPacketByHash(t *testing.T) {
	client := newBscClient(t)
	packets, err := client.GetPacketsByHash("0x7755e0105e5cd796cced4dca65bd11cb9c11160eaab08348cf56f746a28394c1")
	require.NoError(t, err)
	require.NotNil(t, packets.BizPackets)
	var data transfertypes.FungibleTokenPacketData
	err = data.DecodeBytes(packets.BizPackets[0].DataList[0])
	require.NoError(t, err)
	fmt.Println(data.String())
}

func TestGetProofIndex(t *testing.T) {
	rpcClient, err := rpc.DialContext(context.Background(), testUrl)
	bscClient := &Bsc{gethRpcCli: rpcClient}

	require.NoError(t, err)

	for i := int64(205) - 100; i <= 205+100; i++ {
		hash := crypto.Keccak256Hash(
			host.PacketCommitmentKey("bsctest", "teleport", 1),
			common.LeftPadBytes(big.NewInt(i).Bytes(), 32),
		)
		proof, err := bscClient.getProof(context.Background(), common.HexToAddress("0x6cc67656660827f9e4810ed3657b5fdab49a553d"), []string{hexutil.Encode(hash.Bytes())}, big.NewInt(18609747))
		require.NoError(t, err)
		if len(proof.StorageProof) > 1 || proof.StorageProof[0].Value.Uint64() > 0 {
			t.Log(i)
		}
	}
}

func newBscClient(t *testing.T) *Bsc {
	optPrivKey := "FB0536CF27B7F16EAB7F8BBD1771980E83ECE69F50BE30A7161D7E643645958D"

	contractCfgGroup := NewContractCfgGroup()
	contractCfgGroup.Packet.Addr = "0x6cc67656660827f9e4810ed3657b5fdab49a553d"
	contractCfgGroup.Packet.Topic = "PacketSent((uint64,string,string,string,string[],bytes[]))"
	contractCfgGroup.Packet.OptPrivKey = optPrivKey

	contractCfgGroup.AckPacket.Addr = "0x6cc67656660827f9e4810ed3657b5fdab49a553d"
	contractCfgGroup.AckPacket.Topic = "AckWritten((uint64,string,string,string,string[],bytes[]),bytes)"
	contractCfgGroup.AckPacket.OptPrivKey = optPrivKey

	contractCfgGroup.Client.Addr = "0x0053023426adf026c59c80c1880b065b71759dc5"
	contractCfgGroup.Client.Topic = ""
	contractCfgGroup.Client.OptPrivKey = optPrivKey

	contractCfgGroup.Transfer.Addr = "0x8e464b45f4cfb84c5f360d922afc338e56592625"
	contractCfgGroup.Transfer.Topic = "Transfer((string,uint256,string,string))"
	contractCfgGroup.Transfer.OptPrivKey = optPrivKey
	contractBindOptsCfg := NewContractBindOptsCfg()
	contractBindOptsCfg.ChainID = testId
	contractBindOptsCfg.ClientPrivKey = optPrivKey
	contractBindOptsCfg.PacketPrivKey = optPrivKey
	contractBindOptsCfg.GasLimit = 2000000

	chainCfg := NewChainConfig()
	chainCfg.ContractCfgGroup = contractCfgGroup
	chainCfg.ContractBindOptsCfg = contractBindOptsCfg
	chainCfg.ChainType = types.BSC
	chainCfg.ChainName = "BSC"
	chainCfg.ChainURI = testUrl
	chainCfg.ChainID = testId
	chainCfg.Slot = 4
	chainCfg.UpdateClientFrequency = 10

	bscClient, err := newBsc(chainCfg)
	if err != nil {
		t.Fatal(err)
	}
	return bscClient
}
