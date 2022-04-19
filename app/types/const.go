package types

const (
	Tendermint          = "tendermint"
	BSC                 = "bsc"
	ETH                 = "eth"
	EventTypeSendPacket = "xibc.core.packet.v1.EventSendPacket"
	EventTypeWriteAck   = "xibc.core.packet.v1.EventWriteAck"
	CommitmentPoof      = "commitment"
	AckProof            = "ack"
	Ack                 = "ack"
	Packet              = "packet"

	StatusActive  = "active"
	StatusStopped = "stopped"

	RetryTimes = 3

	// ChannelTendermintToTendermint = "tendermint,tendermint"
	// ChannelTendermintToEth        = "tendermint,eth"
	// ChannelEthToTendermint        = "eth,tendermint"
)
