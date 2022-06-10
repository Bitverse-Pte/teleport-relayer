package channels

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"

	"github.com/teleport-network/teleport-relayer/app/types"
	"github.com/teleport-network/teleport-relayer/app/types/errors"
)

func (c *Channel) GetMsg(fromBlock, toBlock uint64) ([]sdk.Msg, error) {
	packets, err := c.chainA.GetPackets(fromBlock, toBlock, c.chainB.ChainType()) // TODO
	if err != nil {
		return nil, err
	}
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return nil, err
	}
	delayHeight, err := c.chainB.GetLightClientDelayHeight(c.chainA.ChainName())
	if err != nil {
		return nil, err
	}
	proofHeight := clientState.GetLatestHeight().GetRevisionHeight() - delayHeight
	var relayPackets []sdk.Msg
	if len(packets.BizPackets) != 0 {
		c.logger.Printf("has queried packet number:%v", len(packets.BizPackets))
		for _, p := range packets.BizPackets {
			c.logger.Printf("packet detail:%v,%v,%v\n", p.SrcChain, p.DstChain, p.Sequence)
		}

	}
	if len(packets.AckPackets) != 0 {
		c.logger.Printf("has queried ack number:%v", len(packets.AckPackets))
		for _, p := range packets.AckPackets {
			c.logger.Printf("packet detail:%v,%v,%v\n", p.Packet.SrcChain, p.Packet.DstChain, p.Packet.Sequence)
		}
	}

	for _, pack := range packets.BizPackets {
		if pack.SrcChain == c.chainB.ChainName() {
			continue
		}
		if err := c.chainA.GetCommitmentsPacket(pack.SrcChain, pack.DstChain, pack.Sequence); err != nil {
			if strings.Contains(err.Error(), "connection") {
				return nil, err
			}
			continue
		}
		// skip receipted
		isNotReceipt, err := c.chainB.GetReceiptPacket(pack.SrcChain, pack.DstChain, pack.Sequence)
		if err != nil {
			return nil, err
		}
		if isNotReceipt {
			c.logger.Printf("packet has been received,sourchain:%v,destchain:%v,sequence:%v", pack.SrcChain, pack.DstChain, pack.Sequence)
			continue
		}
		proof, err := c.chainA.GetProof(pack.SrcChain, pack.DstChain, pack.Sequence, proofHeight, types.CommitmentPoof)
		if err != nil {
			return nil, errors.ErrGetProof
		}
		pac, err := pack.ABIPack()
		if err != nil {
			return nil, errors.ErrGetProof
		}
		recvPacket := &packettypes.MsgRecvPacket{
			Packet:          pac,
			ProofCommitment: proof,
			ProofHeight: clienttypes.Height{
				RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
				RevisionHeight: proofHeight,
			},
		}
		relayPackets = append(relayPackets, recvPacket)
	}
	for _, pack := range packets.AckPackets {
		if err := c.chainB.GetCommitmentsPacket(
			pack.Packet.SrcChain,
			pack.Packet.DstChain,
			pack.Packet.Sequence,
		); err != nil {
			if strings.Contains(err.Error(), "connection") {
				return nil, fmt.Errorf("failed to get commitment packet")
			}
			continue
		}
		// skip receipted
		isNotReceipt, err := c.chainB.GetReceiptPacket(pack.Packet.SrcChain, pack.Packet.DstChain, pack.Packet.Sequence)
		if err != nil {
			return nil, err
		}
		if isNotReceipt {
			c.logger.Printf("ack has been received,sourchain:%v,destchain:%v,sequence:%v", pack.Packet.SrcChain, pack.Packet.DstChain, pack.Packet.Sequence)
			continue
		}
		// query proof
		proof, err := c.chainA.GetProof(
			pack.Packet.SrcChain,
			pack.Packet.DstChain,
			pack.Packet.Sequence,
			proofHeight,
			types.AckProof,
		)
		if err != nil {
			return nil, errors.ErrGetProof
		}
		packetBytes, err := pack.Packet.ABIPack()
		if err != nil {
			return nil, err
		}
		recvPacket := &packettypes.MsgAcknowledgement{
			Packet:          packetBytes,
			Acknowledgement: pack.Acknowledgement,
			ProofAcked:      proof,
			ProofHeight: clienttypes.Height{
				RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
				RevisionHeight: proofHeight,
			},
		}
		relayPackets = append(relayPackets, recvPacket)
	}
	return relayPackets, nil
}

func (c *Channel) GetMsgByHash(hash string) ([]sdk.Msg, error) {
	packets, err := c.chainA.GetPacketsByHash(hash)
	if err != nil {
		return nil, err
	}
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return nil, err
	}
	delayHeight, err := c.chainB.GetLightClientDelayHeight(c.chainA.ChainName())
	if err != nil {
		return nil, err
	}
	proofHeight := clientState.GetLatestHeight().GetRevisionHeight() - delayHeight
	var relayPackets []sdk.Msg
	if len(packets.BizPackets) != 0 {
		c.logger.Printf("has queried packet number:%v", len(packets.BizPackets))
		for _, p := range packets.BizPackets {
			c.logger.Printf("packet detail:%v,%v,%v\n", p.SrcChain, p.DstChain, p.Sequence)
		}

	}
	if len(packets.AckPackets) != 0 {
		c.logger.Printf("has queried ack number:%v", len(packets.AckPackets))
		for _, p := range packets.AckPackets {
			c.logger.Printf("packet detail:%v,%v,%v\n", p.Packet.SrcChain, p.Packet.DstChain, p.Packet.Sequence)
		}
	}

	for _, pack := range packets.BizPackets {
		if pack.SrcChain == c.chainB.ChainName() {
			continue
		}
		if err := c.chainA.GetCommitmentsPacket(pack.SrcChain, pack.DstChain, pack.Sequence); err != nil {
			if strings.Contains(err.Error(), "connection") {
				return nil, err
			}
			continue
		}
		// skip receipted
		isNotReceipt, err := c.chainB.GetReceiptPacket(pack.SrcChain, pack.DstChain, pack.Sequence)
		if err != nil {
			return nil, err
		}
		if isNotReceipt {
			c.logger.Printf("packet has been received,sourchain:%v,destchain:%v,sequence:%v", pack.SrcChain, pack.DstChain, pack.Sequence)
			continue
		}
		proof, err := c.chainA.GetProof(pack.SrcChain, pack.DstChain, pack.Sequence, proofHeight, types.CommitmentPoof)
		if err != nil {
			return nil, errors.ErrGetProof
		}
		packet, err := pack.ABIPack()
		if err != nil {
			return nil, errors.ErrGetProof
		}
		recvPacket := &packettypes.MsgRecvPacket{
			Packet:          packet,
			ProofCommitment: proof,
			ProofHeight: clienttypes.Height{
				RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
				RevisionHeight: proofHeight,
			},
		}
		relayPackets = append(relayPackets, recvPacket)
	}
	for _, pack := range packets.AckPackets {
		if err := c.chainB.GetCommitmentsPacket(
			pack.Packet.SrcChain,
			pack.Packet.DstChain,
			pack.Packet.Sequence,
		); err != nil {
			if strings.Contains(err.Error(), "connection") {
				return nil, fmt.Errorf("failed to get commitment packet")
			}
			continue
		}
		// skip receipted
		isNotReceipt, err := c.chainB.GetReceiptPacket(pack.Packet.SrcChain, pack.Packet.DstChain, pack.Packet.Sequence)
		if err != nil {
			return nil, err
		}
		if isNotReceipt {
			c.logger.Printf("packet has been received,sourchain:%v,destchain:%v,sequence:%v", pack.Packet.SrcChain, pack.Packet.DstChain, pack.Packet.Sequence)
			continue
		}
		// query proof
		proof, err := c.chainA.GetProof(
			pack.Packet.SrcChain,
			pack.Packet.DstChain,
			pack.Packet.Sequence,
			proofHeight,
			types.AckProof,
		)
		if err != nil {
			return nil, errors.ErrGetProof
		}
		packet, err := pack.Packet.ABIPack()
		if err != nil {
			return nil, errors.ErrGetProof
		}
		recvPacket := &packettypes.MsgAcknowledgement{
			Packet:          packet,
			Acknowledgement: pack.Acknowledgement,
			ProofAcked:      proof,
			ProofHeight: clienttypes.Height{
				RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
				RevisionHeight: proofHeight,
			},
		}
		relayPackets = append(relayPackets, recvPacket)
	}
	return relayPackets, nil
}

//
//func (c *Channel) filterPacket(packet *packettypes.Packet) bool {
//	return (packet.SrcChain != c.chainA.ChainName() && packet.RelayChain != c.chainA.ChainName()) ||
//		(packet.DstChain != c.chainB.ChainName() && packet.RelayChain != c.chainB.ChainName()) ||
//		!(packet.RelayChain == c.chainA.ChainName() && packet.RelayChain == c.chainB.ChainName())
//}
