package eth

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/stretchr/testify/require"
	"github.com/teleport-network/teleport/x/xibc/core/host"
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"

	"github.com/teleport-network/teleport-relayer/app/types"
)

const (
	rinkeby   = "https://rinkeby.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	rinkebyID = 4
)

func TestGetProofIndex(t *testing.T) {
	rpcClient, err := rpc.DialContext(context.Background(), rinkeby)
	client := Eth{gethRpcCli: rpcClient}

	require.NoError(t, err)

	for i := int64(205) - 100; i <= 205+100; i++ {
		hash := crypto.Keccak256Hash(
			host.PacketCommitmentKey("rinkeby", "teleport", 1),
			common.LeftPadBytes(big.NewInt(i).Bytes(), 32),
		)
		proof, err := client.getProof(
			context.Background(),
			common.HexToAddress("0xa5ba9eaaa03901870494a6d1f957dd48daec9cf4"),
			[]string{hexutil.Encode(hash.Bytes())},
			big.NewInt(10821006),
		)
		require.NoError(t, err)
		if len(proof.StorageProof) > 1 || proof.StorageProof[0].Value.Uint64() > 0 {
			t.Log(i)
		}
	}
}

func TestGetPacket(t *testing.T) {
	client := getEth(t)

	fromHeight := uint64(10691113)
	toHeight := uint64(10691113)

	for i := fromHeight; i <= toHeight; i++ {
		packets, err := client.GetPackets(i, i, "")
		require.NoError(t, err)
		if len(packets.BizPackets) != 0 {
			fmt.Println(i)
		}
	}

	packets, err := client.GetPackets(fromHeight, toHeight, "")
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
		t.Log(transferData.String())

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
	client := getEth(t)
	packets, err := client.GetPacketsByHash("0x1e110b269f95c5626479f4fd178708dec7006a9c06d59ff86192c69e2bea2644")
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

func getEth(t *testing.T) *Eth {
	optPrivKey := "d10f695d6cbe3d12808a23ba10b5d1fc407dbe0caabb18935e02aedcec8b358b"

	contractCfgGroup := NewContractCfgGroup()
	contractCfgGroup.Packet.Addr = "0xf7268301384fb751e49fafdacd02c693eabb142c"
	contractCfgGroup.Packet.Topic = "PacketSent(bytes)"
	contractCfgGroup.Packet.OptPrivKey = optPrivKey

	contractCfgGroup.AckPacket.Addr = "0xf7268301384fb751e49fafdacd02c693eabb142c"
	contractCfgGroup.AckPacket.Topic = "AckWritten((string,string,uint64,string,bytes,bytes,string,uint64),bytes)"
	contractCfgGroup.AckPacket.OptPrivKey = optPrivKey

	contractCfgGroup.Client.Addr = "0xa46d0b4ed205bf63cd1e2edffef2552b8930c479"
	contractCfgGroup.Client.Topic = ""
	contractCfgGroup.Client.OptPrivKey = optPrivKey

	contractBindOptsCfg := NewContractBindOptsCfg()
	contractBindOptsCfg.ChainID = rinkebyID
	contractBindOptsCfg.ClientPrivKey = optPrivKey
	contractBindOptsCfg.PacketPrivKey = optPrivKey
	contractBindOptsCfg.GasLimit = 2000000

	chainCfg := NewChainConfig()
	chainCfg.ContractCfgGroup = contractCfgGroup
	chainCfg.ContractBindOptsCfg = contractBindOptsCfg
	chainCfg.ChainType = types.ETH
	chainCfg.ChainName = "ETH"
	chainCfg.ChainURI = rinkeby
	chainCfg.ChainID = rinkebyID
	chainCfg.Slot = 4
	chainCfg.UpdateClientFrequency = 10

	bscClient, err := newEth(chainCfg)
	if err != nil {
		t.Fatal(err)
	}
	return bscClient
}
