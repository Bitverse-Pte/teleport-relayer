package channels

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/teleport-network/teleport-relayer/app/types"
	"github.com/teleport-network/teleport-relayer/app/types/errors"
	xibceth "github.com/teleport-network/teleport/x/xibc/clients/light-clients/eth/types"
	xibctendermint "github.com/teleport-network/teleport/x/xibc/clients/light-clients/tendermint/types"
	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"
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
	}
	if len(packets.AckPackets) != 0 {
		c.logger.Printf("has queried ack number:%v", len(packets.AckPackets))
	}

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
			c.logger.Printf("packet has been recived,sourchain:%v,destchain:%v,sequence:%v", pack.SourceChain, pack.DestinationChain, pack.Sequence)
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
				return nil, fmt.Errorf("failed to get commitment packet")
			}
			continue
		}
		// skip receipted
		isNotReceipt, err := c.chainB.GetReceiptPacket(pack.Packet.SourceChain, pack.Packet.DestinationChain, pack.Packet.Sequence)
		if err != nil {
			return nil, err
		}
		if isNotReceipt {
			c.logger.Printf("packet has been recived,sourchain:%v,destchain:%v,sequence:%v", pack.Packet.SourceChain, pack.Packet.DestinationChain, pack.Packet.Sequence)
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

func (c *Channel) filterPacket(packet *packettypes.Packet) bool {
	return (packet.SourceChain != c.chainA.ChainName() && packet.RelayChain != c.chainA.ChainName()) ||
		(packet.DestinationChain != c.chainB.ChainName() && packet.RelayChain != c.chainB.ChainName()) ||
		!(packet.RelayChain == c.chainA.ChainName() && packet.RelayChain == c.chainB.ChainName())
}

func makeCodec() *codec.ProtoCodec {
	ir := codectypes.NewInterfaceRegistry()
	clienttypes.RegisterInterfaces(ir)
	govtypes.RegisterInterfaces(ir)
	xibctendermint.RegisterInterfaces(ir)
	xibceth.RegisterInterfaces(ir)
	packettypes.RegisterInterfaces(ir)
	ir.RegisterInterface("cosmos.v1beta1.Msg", (*sdk.Msg)(nil))
	tx.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)
	return codec.NewProtoCodec(ir)
}
