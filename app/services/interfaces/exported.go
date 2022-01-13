package interfaces

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/teleport-network/teleport/x/xibc/exported"

	"github.com/teleport-network/teleport-relayer/app/services/model"
	"github.com/teleport-network/teleport-relayer/app/types"
)

type IChain interface {
	GetCrossChainPacketsByHeight(height uint64, destChainType string) ([]model.CrossPacket, error)
	GetPackets(height uint64, destChainType string) (*types.Packets, error) // TODO update eth interface
	GetProof(sourChainName, destChainName string, sequence uint64, height uint64, typ string) ([]byte, error)
	RelayPackets(msgs []sdk.Msg) error
	GetCommitmentsPacket(sourChainName, destChainName string, sequence uint64) error
	GetReceiptPacket(sourChainName, destChianName string, sequence uint64) (bool, error)
	GetBlockHeader(*types.GetBlockHeaderReq) (exported.Header, error)
	GetBlockTimestamp(height uint64) (uint64, error)
	GetLightClientState(string) (exported.ClientState, error)
	GetLightClientConsensusState(string, uint64) (exported.ConsensusState, error)
	GetLatestHeight() (uint64, error)
	GetLightClientDelayHeight(string) (uint64, error)
	GetLightClientDelayTime(string) (uint64, error)
	UpdateClient(header exported.Header, chainName string) error
	GetResult(hash string) (uint64, error)
	ChainName() string
	UpdateClientFrequency() uint64
	ChainType() string
}
