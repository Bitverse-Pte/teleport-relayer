package channels

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"

	"github.com/teleport-network/teleport-relayer/app/services/model"
	"github.com/teleport-network/teleport-relayer/app/types"
	"github.com/teleport-network/teleport-relayer/app/types/errors"
)

func (c *Channel) packetSync() error {
	defer func() {
		if err := recover(); err != nil {
			c.logger.Printf("packetSync  panic:%v", err)
		}
	}()
	syncHeight := c.PacketPool.syncHeight
	if syncHeight > c.relayHeight+3 {
		time.Sleep(3 * time.Second)
		return nil
	}
	c.logger.Println("syncHeight:", syncHeight)
	sourceHeight, err := c.chainA.GetLatestHeight()
	if err != nil {
		return err
	}
	if syncHeight > sourceHeight {
		time.Sleep(3 * time.Second)
	}
	packets, err := c.chainA.GetPackets(syncHeight, c.chainB.ChainType()) // TODO
	if err != nil {
		return err
	}
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return err
	}
	var relayPackets []sdk.Msg
	for _, pack := range packets.BizPackets {
		if pack.SourceChain == c.chainB.ChainName() || (pack.DestinationChain != c.chainB.ChainName() && pack.RelayChain != c.chainB.ChainName()) {
			continue
		}
		if err := c.chainA.GetCommitmentsPacket(pack.SourceChain, pack.DestinationChain, pack.Sequence); err != nil {
			if strings.Contains(err.Error(), "connection") {
				return err
			}
			continue
		}
		// 3.2 get receipt packet from chainB chain
		isNotReceipt, err := c.chainB.GetReceiptPacket(pack.SourceChain, pack.DestinationChain, pack.Sequence)
		if err != nil {
			return err
		}
		// if receipt exist, skip
		if isNotReceipt {
			continue
		}
		proof, err := c.chainA.GetProof(pack.SourceChain, pack.DestinationChain, pack.Sequence, syncHeight+1, types.CommitmentPoof)
		if err != nil {
			return errors.ErrGetProof
		}
		recvPacket := &packettypes.MsgRecvPacket{
			Packet:          pack,
			ProofCommitment: proof,
			ProofHeight: clienttypes.Height{
				RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
				RevisionHeight: syncHeight + 1,
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
				return errors.ErrGetCommitmentPacket
			}
			continue
		}
		// query proof
		proof, err := c.chainA.GetProof(
			pack.Packet.SourceChain,
			pack.Packet.DestinationChain,
			pack.Packet.Sequence,
			c.relayHeight,
			types.AckProof,
		)
		if err != nil {
			return errors.ErrGetProof
		}
		recvPacket := &packettypes.MsgAcknowledgement{
			Packet:          pack.Packet,
			Acknowledgement: pack.Acknowledgement,
			ProofAcked:      proof,
			ProofHeight: clienttypes.Height{
				RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
				RevisionHeight: c.relayHeight,
			},
		}
		relayPackets = append(relayPackets, recvPacket)
	}
	if len(relayPackets) != 0 {
		c.PacketPool.Write(syncHeight, relayPackets)
	}
	c.PacketPool.syncHeight++
	return nil
}

func (c *Channel) DeleteRelayedPacket() {
	for {
		// TODO
		for height := range c.PacketPool.BP {
			if height < c.relayHeight {
				c.logger.Printf("block packet Height %v < relay height %v\n ", height, c.relayHeight)
				c.PacketPool.Delete(height)
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func (c *Channel) packetSyncToDB() error {
	syncHeight := c.PacketPool.syncHeight
	if syncHeight > c.relayHeight+3 {
		time.Sleep(3 * time.Second)
		return nil
	}
	c.logger.Println("syncHeight:", syncHeight)
	sourceHeight, err := c.chainA.GetLatestHeight()
	if err != nil {
		return err
	}
	if syncHeight > sourceHeight {
		time.Sleep(3 * time.Second)
	}
	crossChainPackets, err := c.chainA.GetCrossChainPacketsByHeight(syncHeight, c.chainB.ChainType()) // TODO
	if err != nil {
		return err
	}
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return err
	}

	return c.PacketDBPool.DB.Client.Transaction(
		func(tx *gorm.DB) error {
			for _, ccpt := range crossChainPackets {
				if ccpt.Packet != nil {
					proof, err := c.chainA.GetProof(ccpt.Packet.SourceChain, ccpt.Packet.DestinationChain, ccpt.Packet.Sequence, syncHeight+1, types.CommitmentPoof)
					if err != nil {
						return errors.ErrGetProof
					}
					recvPacket := &packettypes.MsgRecvPacket{
						Packet:          *ccpt.Packet,
						ProofCommitment: proof,
						ProofHeight: clienttypes.Height{
							RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
							RevisionHeight: syncHeight + 1,
						},
					}
					recvPacketByte, err := json.Marshal(recvPacket)
					if err != nil {
						return err
					}
					crossChainPacket := model.CrossChainPacket{
						Commitment: ccpt.Commitment,
						Packet:     string(recvPacketByte),
						AckPacket:  "",
						Height:     syncHeight,
						TxHash:     ccpt.TxHash,
						Sender:     ccpt.Sender,
						Status:     ccpt.Status,
					}
					if err := createOrUpdate(tx, crossChainPacket); err != nil {
						return err
					}
				}
				if ccpt.AckPacket != nil {
					proof, err := c.chainA.GetProof(ccpt.AckPacket.Packet.SourceChain, ccpt.AckPacket.Packet.DestinationChain, ccpt.AckPacket.Packet.Sequence, syncHeight+1, types.CommitmentPoof)
					if err != nil {
						return errors.ErrGetProof
					}
					recvPacket := &packettypes.MsgAcknowledgement{
						Packet:          ccpt.AckPacket.Packet,
						Acknowledgement: ccpt.AckPacket.Acknowledgement,
						ProofAcked:      proof,
						ProofHeight: clienttypes.Height{
							RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
							RevisionHeight: c.relayHeight,
						},
					}
					recvPacketByte, err := json.Marshal(recvPacket)
					if err != nil {
						return err
					}
					crossChainPacket := model.CrossChainPacket{
						Commitment: ccpt.Commitment,
						Packet:     string(recvPacketByte),
						AckPacket:  "",
						Height:     syncHeight,
						TxHash:     ccpt.TxHash,
						Sender:     ccpt.Sender,
						Status:     ccpt.Status,
					}
					if err := createOrUpdate(tx, crossChainPacket); err != nil {
						return err
					}
				}
			}
			return nil
		},
	)
}

func createOrUpdate(tx *gorm.DB, data interface{}) error {
	if tx.Model(&model.CrossChainPacket{}).Select(&data).RecordNotFound() {
		return tx.Create(&data).Error
	}
	return tx.Update(&data).Error
}
