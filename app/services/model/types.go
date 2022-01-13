package model

import (
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"

	"github.com/teleport-network/teleport-relayer/app/types"
)

type CrossChainPacket struct {
	Commitment string
	Packet     string
	AckPacket  string
	Height     uint64
	TxHash     string
	Sender     string
	Status     int8
}

type CrossPacket struct {
	Commitment string
	Packet     *packettypes.Packet
	AckPacket  *types.AckPacket
	Height     uint64
	TxHash     string
	Sender     string
	Status     int8
}

type CrossChainPacketStatus string

const (
	PacketReceived CrossChainPacketStatus = "packetReceived"
	AckReceived    CrossChainPacketStatus = "ackReceived"
)
