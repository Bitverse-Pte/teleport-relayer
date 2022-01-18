package bsc

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/teleport-network/teleport-relayer/app/services/interfaces"
	"github.com/teleport-network/teleport-relayer/app/types"
)

const (
	testUrl = "https://data-seed-prebsc-1-s1.binance.org:8545/"
	testId  = 97
	epoch   = uint64(200)
)

func TestNewBsc(t *testing.T) {
	bscClient := newBscClient(t)

	latestHeight, err := bscClient.GetLatestHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(latestHeight)
	decodeString, err := hex.DecodeString("0a0f74656c65706f72745f393030302d311204080110031a051080ac8f04220510808ac4042a02100a32040801100a42060a0478696263")
	if err != nil {
		return
	}
	fmt.Println(string(decodeString))
}

func TestGetStatesJsons(t *testing.T) {
	bscClient := newBscClient(t)
	latestHeight, err := bscClient.GetLatestHeight()
	if err != nil {
		t.Fatal(err)
	}

	//todo gen Jsons
	startHeight := latestHeight - latestHeight%epoch - epoch
	t.Log("bsc height = ", startHeight)
}

func newBscClient(t *testing.T) interfaces.IChain {
	optPrivKey := "4f706b587618e242f45f9f67fb5cbb290902c7ff5828c468ee53138ef8a26945"

	contractCfgGroup := NewContractCfgGroup()
	contractCfgGroup.Packet.Addr = "0x2A212D09038c848A0d79a42E0Ab88B5FD8B1AD85"
	contractCfgGroup.Packet.Topic = "PacketSent((uint64,string,string,string,string,bytes))"
	contractCfgGroup.Packet.OptPrivKey = optPrivKey

	contractCfgGroup.AckPacket.Addr = "0x2A212D09038c848A0d79a42E0Ab88B5FD8B1AD85"
	contractCfgGroup.AckPacket.Topic = "AckWritten((uint64,string,string,string,string,bytes),bytes)"
	contractCfgGroup.AckPacket.OptPrivKey = optPrivKey

	contractCfgGroup.Client.Addr = "0x53176d71Ac1AD08cF5a7e54aF1EdF5657B2419eC"
	contractCfgGroup.Client.Topic = ""
	contractCfgGroup.Client.OptPrivKey = optPrivKey

	contractCfgGroup.Transfer.Addr = "0xD002C2fC0C1c0883F85eA1aa0305c7Fd7CD829e0"
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
