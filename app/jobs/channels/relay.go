package channels

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sync"
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
	if (height+delayHeight < c.clientHeight) || (c.chainA.ChainType() == types.Tendermint && height < chainAHeight) {
		return nil
	}
	return fmt.Errorf("need wait client update to height %d ! clientHeight=%v < relayHeighta:%v", height+delayHeight, c.clientHeight, height+delayHeight)
}

func (c *Channel) checkClient() error {
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return fmt.Errorf("get lightClient client state fail:%+v", err)
	}
	if clientState.GetLatestHeight().GetRevisionHeight() == c.checkHeight {
		panic("client height not updated for a long time")
	} else {
		c.checkHeight = clientState.GetLatestHeight().GetRevisionHeight()
	}
	return nil
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
	// update larger height
	reqHeight := updateHeight
	if updateHeight < chainAHeight {
		reqHeight = chainAHeight
	}
	revisionNumber := clientState.GetLatestHeight().GetRevisionNumber()
	var header exported.Header
	req := &types.GetBlockHeaderReq{
		LatestHeight:   reqHeight,
		TrustedHeight:  revisionHeight,
		RevisionNumber: revisionNumber,
	}
	header, err = c.chainA.GetBlockHeader(req)
	if err != nil {
		return err
	}
	return c.chainB.UpdateClient(header, c.chainA.ChainName())
}

func (c *Channel) batchGetBlockHeader(reqHeight, revisionHeight, revisionNumber uint64) ([]exported.Header, error) {
	times := 5
	headers := make([]exported.Header, times)
	var l sync.Mutex
	var wg sync.WaitGroup
	wg.Add(times)
	for i := reqHeight; i < reqHeight+uint64(times); i++ {
		go func(height uint64) {
			defer wg.Done()
			var header exported.Header
			var err error
			req := &types.GetBlockHeaderReq{
				LatestHeight:   height,
				TrustedHeight:  revisionHeight,
				RevisionNumber: revisionNumber,
			}
			header, err = c.chainA.GetBlockHeader(req)
			if err != nil {
				c.logger.Errorf("GetBlockHeader error:%+v", err)
				return
			}
			l.Lock()
			headers[height-reqHeight] = header
			l.Unlock()
		}(i)
	}
	wg.Wait()
	for _, h := range headers {
		if h == nil {
			return nil, fmt.Errorf("get headers failed")
		}
	}
	return headers, nil
}

func (c *Channel) RelayTask(s *gocron.Scheduler) {
	relayJobs, err := s.Every(int(c.relayFrequency)).Seconds().Do(func() {
		time.Sleep(time.Duration(c.extraWait*c.relayFrequency) * time.Second)
		c.logger.Infof("start relay %+v! height : %+v", c.chainA.ChainName(), c.relayHeight)
		c.UpdateHeight()
		if err := c.RelayPackets(c.relayHeight); err != nil {
			c.logger.Errorf("RelayPackets err : %+v", err)
			return
		}
	})
	if err != nil {
		panic(err)
	}
	relayJobs.SingletonMode()
	if c.chainA.ChainType() == types.ETH || c.chainA.ChainType() == types.BSC {
		updateJobs, err := s.Every(int(c.relayFrequency)).Seconds().Do(func() {
			time.Sleep(time.Duration(c.extraWait*c.relayFrequency) * time.Second)
			if err := c.evmClientUpdate(); err != nil {
				c.logger.Errorf("EvmClientUpdate err : %+v", err)
				return
			}
		})
		if err != nil {
			panic(err)
		}
		updateJobs.SingletonMode()
		checkJob, err := s.Every(10).Minute().Do(func() {
			if err := c.checkClient(); err != nil {
				c.logger.Errorf("checkClient err : %+v", err)
			}
		})
		if err != nil {
			panic(err)
		}
		checkJob.SingletonMode()
	}
}

func (c *Channel) RelayPackets(height uint64) error {
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
	updateHeight := height
	var pkt []sdk.Msg
	verifyHeight := c.clientHeight
	if c.chainA.ChainType() == types.Tendermint {
		verifyHeight = chainAHeight
	}
	if height+delayHeight+10 < verifyHeight {
		pkt, err = c.GetMsg(height, height+9)
		if err != nil {
			return fmt.Errorf("batchGetMsg error:%+v", err)
		}
		updateHeight += 10
	} else if height+delayHeight < verifyHeight {
		pkt, err = c.GetMsg(height, height)
		if err != nil {
			return fmt.Errorf("get msg err:%+v", err)
		}
		updateHeight += 1
	} else {
		time.Sleep(10 * time.Second)
		return fmt.Errorf("height + delayHeight >= verifyHeight")
	}
	if len(pkt) == 0 {
		c.relayHeight = updateHeight
		return nil
	}
	if c.chainA.ChainType() == types.Tendermint {
		if err := c.UpdateClientByHeight(updateHeight); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	if err := c.chainB.RelayPackets(pkt); err != nil {
		if !handleRecvPacketsError(err) {
			return fmt.Errorf("failed to recv packet:%v", err.Error())
		}
	}
	c.relayHeight = updateHeight
	return nil
}

