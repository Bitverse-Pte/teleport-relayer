package channels

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"

	"github.com/teleport-network/teleport-relayer/app/types"
	"github.com/teleport-network/teleport-relayer/app/types/errors"
)

func (c *Channel) GetMsg(height uint64) ([]sdk.Msg, error) {
	packets, err := c.chainA.GetPackets(height, c.chainB.ChainType()) // TODO
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
	if proofHeight < height+1 {
		proofHeight = height + 1
	}
	var relayPackets []sdk.Msg
	for _, pack := range packets.BizPackets {
		if pack.SourceChain == c.chainB.ChainName() || (pack.DestinationChain != c.chainB.ChainName() && pack.RelayChain != c.chainB.ChainName()) {
			continue
		}
		if err := c.chainA.GetCommitmentsPacket(pack.SourceChain, pack.DestinationChain, pack.Sequence); err != nil {
			if strings.Contains(err.Error(), "connection") {
				return nil, err
			}
			continue
		}
		// skip receipted
		isNotReceipt, err := c.chainB.GetReceiptPacket(pack.SourceChain, pack.DestinationChain, pack.Sequence)
		if err != nil {
			return nil, err
		}
		if isNotReceipt {
			continue
		}
		proof, err := c.chainA.GetProof(pack.SourceChain, pack.DestinationChain, pack.Sequence, proofHeight, types.CommitmentPoof)
		if err != nil {
			return nil, errors.ErrGetProof
		}
		recvPacket := &packettypes.MsgRecvPacket{
			Packet:          pack,
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
			pack.Packet.SourceChain,
			pack.Packet.DestinationChain,
			pack.Packet.Sequence,
		); err != nil {
			if strings.Contains(err.Error(), "connection") {
				return nil, errors.ErrGetCommitmentPacket
			}
			continue
		}
		// skip receipted
		isNotReceipt, err := c.chainB.GetReceiptPacket(pack.Packet.SourceChain, pack.Packet.DestinationChain, pack.Packet.Sequence)
		if err != nil {
			return nil, err
		}
		if isNotReceipt {
			continue
		}
		// query proof
		proof, err := c.chainA.GetProof(
			pack.Packet.SourceChain,
			pack.Packet.DestinationChain,
			pack.Packet.Sequence,
			proofHeight,
			types.AckProof,
		)
		if err != nil {
			return nil, errors.ErrGetProof
		}
		recvPacket := &packettypes.MsgAcknowledgement{
			Packet:          pack.Packet,
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
