package eth

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/teleport-network/teleport-relayer/app/services/eth/contracts/transfer"
	"github.com/teleport-network/teleport-relayer/app/types"
)

func TestNewEth(t *testing.T) {
	url := "https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	optPrivKey := "4f706b587618e242f45f9f67fb5cbb290902c7ff5828c468ee53138ef8a26945"
	var chainID uint64 = 4

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

	// address tokenAddress;
	//        string receiver;
	//        uint256 amount;
	//        string destChain;
	//        string relayChain;

	// struct Packet {
	//        uint64 sequence;
	//        string port;
	//        string sourceChain;
	//        string destChain;
	//        string relayChain;
	//        bytes data;
	//    }

	contractBindOptsCfg := NewContractBindOptsCfg()
	contractBindOptsCfg.ChainID = chainID
	contractBindOptsCfg.ClientPrivKey = optPrivKey
	contractBindOptsCfg.PacketPrivKey = optPrivKey
	contractBindOptsCfg.GasLimit = 2000000
	//contractBindOptsCfg.GasPrice = 1500000000

	//  ropsten: {
	//            url: 'https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161',
	//            gasPrice: 9000000000,
	//            chainId: 3,
	//            gas: 4100000,
	//            accounts: ['4f706b587618e242f45f9f67fb5cbb290902c7ff5828c468ee53138ef8a26945'],
	//        },

	chainCfg := NewChainConfig()
	chainCfg.ContractCfgGroup = contractCfgGroup
	chainCfg.ContractBindOptsCfg = contractBindOptsCfg
	chainCfg.ChainType = types.ETH
	chainCfg.ChainName = "ETH"
	chainCfg.ChainURI = url
	chainCfg.ChainID = chainID
	chainCfg.Slot = 4
	chainCfg.UpdateClientFrequency = 10

	ethClient, err := newEth(chainCfg)
	if err != nil {
		t.Fatal(err)
	}
	latestHeight, err := ethClient.GetLatestHeight()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(latestHeight)
	res, err := ethClient.GetLightClientState("teleport")
	if err != nil {
		fmt.Printf("GetLightClientState ERROR:%v", err)
	}
	fmt.Println(res, err)

	tokenAddress := "0x582e0992cb1EaE9B1AbcBF889EE640626453259F"
	receiver := "0xFd805Fc7f5B60849dbA893168708AAFDD181fCf3"
	destChain := "teleport"
	relayChain := ""
	transferData := transfer.TransferDataTypesERC20TransferData{
		TokenAddress: common.HexToAddress(tokenAddress),
		Receiver:     receiver,
		Amount:       sdk.NewInt(100).BigInt(),
		DestChain:    destChain,
		RelayChain:   relayChain,
	}
	if err := ethClient.TransferERC20(transferData); err != nil {
		fmt.Printf("TransferERC20 ERROR:%v", err)
	}
}

func Test_Hex(t *testing.T) {
	str := "0000000000000000000000000000000000000000000000000000000000000003"
	dataBytes := common.HexToHash(str)
	args := abi.Arguments{
		abi.Argument{Type: Uint64},
	}

	headerBytes, err := args.Unpack(dataBytes.Bytes())
	if err != nil {
		return
	}
	fmt.Println("headerBytes: ", headerBytes)
}

func TestMakeBytes(t *testing.T) {
	// TODO
}
