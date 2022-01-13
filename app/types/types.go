package types

import (
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"
)

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
