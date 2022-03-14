package bsc

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"

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

	packets, err := bscClient.GetPackets(17547678, 17547678, "")
	require.NoError(t, err)
	require.NotNil(t, packets.BizPackets)
	var data transfertypes.FungibleTokenPacketData
	err = data.DecodeBytes(packets.BizPackets[0].DataList[0])
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestGetProofIndex(t *testing.T) {
	rpcClient, err := rpc.DialContext(context.Background(), testUrl)
	bscClient := &Bsc{gethRpcCli: rpcClient}

	require.NoError(t, err)

	for i := int64(155) - 55; i <= 155+100; i++ {
		hash := crypto.Keccak256Hash(
			host.PacketCommitmentKey("bsc-testnet", "teleport", 1),
			common.LeftPadBytes(big.NewInt(i).Bytes(), 32),
		)
		proof, err := bscClient.getProof(context.Background(), common.HexToAddress("0x73f9b1905473e33a30c4b339d3ce87d95a3bfe73"), []string{hexutil.Encode(hash.Bytes())}, big.NewInt(17459374))
		require.NoError(t, err)
		if len(proof.StorageProof) > 1 || proof.StorageProof[0].Value.Uint64() > 0 {
			t.Log(i)
		}
	}
}

func newBscClient(t *testing.T) *Bsc {
	optPrivKey := "FB0536CF27B7F16EAB7F8BBD1771980E83ECE69F50BE30A7161D7E643645958D"

	contractCfgGroup := NewContractCfgGroup()
	contractCfgGroup.Packet.Addr = "0xA6a9AAB1c5E65e1a69cCCF59155ABaA0A555D955"
	contractCfgGroup.Packet.Topic = "PacketSent((uint64,string,string,string,string[],bytes[]))"
	contractCfgGroup.Packet.OptPrivKey = optPrivKey

	contractCfgGroup.AckPacket.Addr = "0xA6a9AAB1c5E65e1a69cCCF59155ABaA0A555D955"
	contractCfgGroup.AckPacket.Topic = "AckWritten((uint64,string,string,string,string[],bytes[]),bytes)"
	contractCfgGroup.AckPacket.OptPrivKey = optPrivKey

	contractCfgGroup.Client.Addr = "0x1012978EDB55F4eD2faEf5CE09cd64965AC38d17"
	contractCfgGroup.Client.Topic = ""
	contractCfgGroup.Client.OptPrivKey = optPrivKey

	contractCfgGroup.Transfer.Addr = "0x1b49147aB0099B8dc03d4a22B15EeAa9403Fa3ED"
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
