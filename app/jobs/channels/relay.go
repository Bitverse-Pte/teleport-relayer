package channels

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"

	"github.com/teleport-network/teleport/x/xibc/exported"

	"github.com/teleport-network/teleport-relayer/app/types"
)

func (c *Channel) validatePacketHeight(height uint64) error {
	delayHeight, err := c.chainB.GetLightClientDelayHeight(c.chainA.ChainName())
	if err != nil {
		return fmt.Errorf("get lightClient delay height fail:%+v", err)

	}
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return fmt.Errorf("get lightClient client state fail:%+v", err)

	}
	c.clientHeight = clientState.GetLatestHeight().GetRevisionHeight()
	chainAHeight, err := c.chainA.GetLatestHeight()
	if err != nil {
		return fmt.Errorf("get latest height error %+v", err)
	}
	if height+delayHeight < c.clientHeight || (c.chainA.ChainType() == types.Tendermint && height < chainAHeight) {
		return nil
	}
	return fmt.Errorf("need wait client update to height %d ! clientHeight=%v < relayHeight=%v", height+delayHeight, c.clientHeight, height+delayHeight)
}

func (c *Channel) UpdateClientByHeight(height uint64) error {
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return err
	}
	revisionHeight := clientState.GetLatestHeight().GetRevisionHeight()
	// 3. Get the latest block currently scanned, and then update
	updateHeight := height + 1
	if c.chainA.ChainType() == types.Tendermint && revisionHeight >= updateHeight {
		c.logger.Println("no need update client")
		return nil
	}
	chainAHeight, err := c.chainA.GetLatestHeight()
	if err != nil {
		return err
	}
	// latest height >= packet height + 1
	if c.chainA.ChainType() == types.Tendermint && clientState.GetLatestHeight().GetRevisionHeight() >= updateHeight {
		c.logger.Println("no need update client")
		return nil
	}
	// update larger height
	reqHeight := updateHeight
	if updateHeight < chainAHeight {
		reqHeight = chainAHeight
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

func (c *Channel) RelayTask(s *gocron.Scheduler) {
	s.Every(5).Seconds().Do(func() {
		c.logger.Infof("start relay %+v! height : %+v", c.chainA.ChainName(), c.relayHeight)
		c.UpdateHeight()
		if err := c.RelayPackets(c.relayHeight); err != nil {
			c.logger.Errorf("RelayPackets err : %+v", err)
			return
		}
		c.relayHeight++
	})
	if c.chainA.ChainType() == types.ETH || c.chainA.ChainType() == types.BSC {
		s.Every(5).Seconds().Do(func() {
			if err := c.EvmClientUpdate(); err != nil {
				c.logger.Errorf("EvmClientUpdate err : %+v", err)
				return
			}
		})
	}
}

func (c *Channel) RelayPackets(height uint64) error {
	if err := c.validatePacketHeight(height); err != nil {
		time.Sleep(20 * time.Second)
		return err
	}
	pkt, err := c.GetMsg(height)
	if err != nil {
		return fmt.Errorf("get msg err:%+v",err)
	}
	if len(pkt) == 0 {
		return nil
	}
	if c.chainA.ChainType() == types.Tendermint {
		if err := c.UpdateClientByHeight(height); err != nil {
			return err
		}
	}
	if err := c.chainB.RelayPackets(pkt); err != nil {
		if !handleRecvPacketsError(err) {
			return fmt.Errorf("failed to recv packet:%v", err.Error())
		}
	}
	return nil
}
