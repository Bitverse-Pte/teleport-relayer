package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"
)

type RpcTransaction struct {
	Tx *ethtypes.Transaction
	TxExtraInfo
}

type TxExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

type GetBlockHeaderReq struct {
	LatestHeight   uint64
	TrustedHeight  uint64
	RevisionNumber uint64
}

type Packets struct {
	BizPackets []packettypes.Packet
	AckPackets []AckPacket
}

type AckPacket struct {
	Packet          packettypes.Packet
	Acknowledgement []byte
}

type ResultTx struct {
	GasWanted int64  `json:"gas_wanted"`
	GasUsed   int64  `json:"gas_used"`
	Hash      string `json:"hash"`
	Height    int64  `json:"height"`
}

type PacketDetail struct {
	ChainName  string
	Sequence   uint64
	SrcChain   string
	DestChain  string
	RelayChain string
	FromHeight uint64
	ToHeight   uint64
	Type       string
	ErrMsg     string
}

func NewPacketDetail(chainName string, sequence uint64, srcChain string, destChain string, relayChain string, fromHeight uint64, toHeight uint64, Ty string) *PacketDetail {
	return &PacketDetail{ChainName: chainName, Sequence: sequence, SrcChain: srcChain, DestChain: destChain, RelayChain: relayChain, FromHeight: fromHeight, ToHeight: toHeight, Type: Ty}
}

func GetPacketDetail(pkt sdk.Msg) PacketDetail {
	switch p := pkt.(type) {
	case *packettypes.MsgRecvPacket:
		return PacketDetail{
			Sequence:   p.Packet.Sequence,
			SrcChain:   p.Packet.SourceChain,
			DestChain:  p.Packet.DestinationChain,
			RelayChain: p.Packet.RelayChain,
			Type:       "Packet",
		}
	case *packettypes.MsgAcknowledgement:
		return PacketDetail{
			Sequence:   p.Packet.Sequence,
			SrcChain:   p.Packet.SourceChain,
			DestChain:  p.Packet.DestinationChain,
			RelayChain: p.Packet.RelayChain,
			Type:       "Acknowledgement",
		}
	}
	return PacketDetail{}
}

func (pd PacketDetail) Equal(srcChain, destChain, relayChain string, sequence uint64) bool {
	return pd.Sequence == sequence && pd.SrcChain == srcChain && pd.DestChain == destChain && pd.RelayChain == relayChain
}
