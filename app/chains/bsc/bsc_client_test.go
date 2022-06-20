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

	"github.com/teleport-network/teleport-relayer/app/types"
	"github.com/teleport-network/teleport/x/xibc/core/host"
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"
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

	fromHeight := uint64(19316505)
	toHeight := uint64(19316555)

	for i := fromHeight; i <= toHeight; i++ {
		packets, err := bscClient.GetPackets(i, i, "")
		require.NoError(t, err)
		if len(packets.BizPackets) != 0 {
			fmt.Println(i)
		}
	}

	packets, err := bscClient.GetPackets(fromHeight, toHeight, "")
	require.NoError(t, err)
	require.NotNil(t, packets.BizPackets)

	for _, v := range packets.BizPackets {
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

}

func TestGetPacketByHash(t *testing.T) {
	client := newBscClient(t)
	packets, err := client.GetPacketsByHash("")
	require.NoError(t, err)
	require.NotNil(t, packets.BizPackets)

	for _, v := range packets.BizPackets {
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
	contractCfgGroup.Packet.Addr = "0x8e84ef5d13a129183b838d833e4ac14eb0c5ceab"
	contractCfgGroup.Packet.Topic = "PacketSent(bytes)"
	contractCfgGroup.Packet.OptPrivKey = optPrivKey

	contractCfgGroup.AckPacket.Addr = "0x8e84ef5d13a129183b838d833e4ac14eb0c5ceab"
	contractCfgGroup.AckPacket.Topic = "AckWritten((string,string,uint64,string,bytes,bytes,string,uint64),bytes)"
	contractCfgGroup.AckPacket.OptPrivKey = optPrivKey

	contractCfgGroup.Client.Addr = "0x4a6a5fbe99259fccf2a4a707f6b6f77a0e80fc97"
	contractCfgGroup.Client.Topic = ""
	contractCfgGroup.Client.OptPrivKey = optPrivKey

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
