package channels

import (
	"time"

	"github.com/teleport-network/teleport/x/xibc/exported"

	"github.com/teleport-network/teleport-relayer/app/types"
	"github.com/teleport-network/teleport-relayer/app/types/errors"
)

func (c *Channel) relay() error {
	defer func() {
		if err := recover(); err != nil {
			c.logger.Printf("relayer panic:%v", err)
		}
	}()
	c.logger.Printf("src chainName = %v\n", c.chainName)
	c.logger.Printf("syncHeight = %v\n", c.PacketPool.syncHeight)
	c.logger.Printf("relayer Height = %v\n", c.PacketPool.syncHeight)
	c.logger.Printf("packet pool length:%v", len(c.PacketPool.BP))
	if !c.validatePacketHeight(c.relayHeight) && c.chainA.ChainType() == types.ETH {
		time.Sleep(3 * time.Second) //TODO
		return nil
	}
	if c.PacketPool.syncHeight < c.relayHeight {
		time.Sleep(3 * time.Second) //TODO
		return nil
	}
	pkt, ok := c.PacketPool.BP[c.relayHeight]
	if !ok {
		if c.relayHeight < c.PacketPool.syncHeight {
			c.relayHeight++
			return nil
		}
		time.Sleep(3 * time.Second)
		return nil
	}
	if c.chainA.ChainType() == types.Tendermint {
		if err := c.UpdateClient(); err != nil {
			return err
		}
	}
	if err := c.chainB.RelayPackets(pkt); err != nil {
		if !handleRecvPacketsError(err) {
			return errors.ErrRecvPacket
		}
	}
	c.relayHeight++
	return nil
}

func (c *Channel) validatePacketHeight(height uint64) bool {
	delayHeight, err := c.chainB.GetLightClientDelayHeight(c.chainA.ChainName())
	if err != nil {
		return false
	}
	if height+delayHeight < c.clientHeight {
		return true
	}
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return false
	}
	c.clientHeight = clientState.GetLatestHeight().GetRevisionHeight()
	return height+delayHeight < c.clientHeight
}

func (c *Channel) UpdateClient() error {
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return err
	}
	revisionHeight := clientState.GetLatestHeight().GetRevisionHeight()
	// 3. Get the latest block currently scanned, and then update
	relayHeight := c.relayHeight + 1

	if relayHeight <= revisionHeight {
		return nil
	}
	chainAHeight, err := c.chainA.GetLatestHeight()
	if err != nil {
		return err
	}
	reqHeight := chainAHeight
	if relayHeight < chainAHeight {
		reqHeight = relayHeight
	}
	if c.chainA.ChainType() == types.Tendermint && clientState.GetLatestHeight().GetRevisionHeight() >= reqHeight {
		c.logger.Println("no need update client")
		return nil
	}
	var header exported.Header
	req := &types.GetBlockHeaderReq{
		LatestHeight:   reqHeight,
		TrustedHeight:  clientState.GetLatestHeight().GetRevisionHeight(),
		RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
	}
	header, err = c.chainA.GetBlockHeader(req)
	if err != nil {
		return err
	}
	return c.chainB.UpdateClient(header, c.chainA.ChainName())
}
