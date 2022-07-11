package types

import (
	"fmt"

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

type SuccessRelay struct {
	ChainName   string
	Sequence    uint64
	SrcChain    string
	DestChain   string
	ReceiveHash string
	Type        string
}

func NewSuccessRelay(chainName string, sequence uint64, srcChain string, destChain string, receiveHash string, Ty string) *SuccessRelay {
	return &SuccessRelay{ChainName: chainName, Sequence: sequence, SrcChain: srcChain, DestChain: destChain, ReceiveHash: receiveHash, Type: Ty}
}

type PacketDetail struct {
	ChainName  string
	Sequence   uint64
	SrcChain   string
	DestChain  string
	FromHeight uint64
	ToHeight   uint64
	Hash       string
	Type       string
	ErrMsg     string
}

func NewPacketDetail(chainName string, sequence uint64, srcChain string, destChain string, fromHeight, toHeight uint64, hash string, Ty string) *PacketDetail {
	return &PacketDetail{ChainName: chainName, Sequence: sequence, SrcChain: srcChain, DestChain: destChain, FromHeight: fromHeight, ToHeight: toHeight, Hash: hash, Type: Ty}
}

func (p *PacketDetail) ToSuccessRelay(receiveHash string) *SuccessRelay {
	return &SuccessRelay{ChainName: p.ChainName, Sequence: p.Sequence, SrcChain: p.SrcChain, DestChain: p.DestChain, ReceiveHash: receiveHash, Type: p.Type}
}

func (p *PacketDetail) PacString() string {
	return fmt.Sprintf("SrcChain: %v ,DestChain: %v ,Sequence: %v ", p.SrcChain, p.DestChain, p.Sequence)
}

func GetPacketDetail(pkt sdk.Msg) PacketDetail {
	switch p := pkt.(type) {
	case *packettypes.MsgRecvPacket:
		var pac packettypes.Packet
		err := pac.ABIDecode(p.Packet)
		if err != nil {
			return PacketDetail{}
		}
		return PacketDetail{
			Sequence:  pac.Sequence,
			SrcChain:  pac.SrcChain,
			DestChain: pac.DstChain,
			Type:      "Packet",
		}
	case *packettypes.MsgAcknowledgement:
		var pac packettypes.Packet
		err := pac.ABIDecode(p.Packet)
		if err != nil {
			return PacketDetail{}
		}
		return PacketDetail{
			Sequence:  pac.Sequence,
			SrcChain:  pac.SrcChain,
			DestChain: pac.DstChain,
			Type:      "Acknowledgement",
		}
	}
	return PacketDetail{}
}

func (pd PacketDetail) Equal(srcChain, destChain string, sequence uint64) bool {
	return pd.Sequence == sequence && pd.SrcChain == srcChain && pd.DestChain == destChain
}
