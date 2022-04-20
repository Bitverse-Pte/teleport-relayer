package interfaces

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/teleport-network/teleport/x/xibc/exported"

	"github.com/teleport-network/teleport-relayer/app/types"
)

type IChain interface {
	GetPackets(fromBlock, toBlock uint64, destChainType string) (*types.Packets, error) // TODO update eth interface
	GetProof(sourChainName, destChainName string, sequence uint64, height uint64, typ string) ([]byte, error)
	RelayPackets(msg sdk.Msg) (string, error)
	GetPacketsByHash(hash string) (*types.Packets, error)
	GetCommitmentsPacket(sourChainName, destChainName string, sequence uint64) error
	GetReceiptPacket(sourChainName, destChainName string, sequence uint64) (bool, error)
	GetBlockHeader(*types.GetBlockHeaderReq) (exported.Header, error)
	GetLightClientState(string) (exported.ClientState, error)
	GetLatestHeight() (uint64, error)
	GetLightClientDelayHeight(string) (uint64, error)
	UpdateClient(header exported.Header, chainName string) error
	BatchUpdateClient(headers []exported.Header, chainName string) error
	GetResult(hash string) (uint64, error)
	ChainName() string
	UpdateClientFrequency() uint64
	ChainType() string
}
