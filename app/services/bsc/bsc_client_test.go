package bsc

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/teleport-network/teleport-relayer/app/services/interfaces"
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

	packets, err := bscClient.GetPackets(15965605, "")
	require.NoError(t, err)
	require.NotNil(t, packets.BizPackets)
}

func newBscClient(t *testing.T) interfaces.IChain {
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
