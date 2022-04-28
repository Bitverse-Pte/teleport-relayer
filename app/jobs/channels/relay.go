package channels

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/teleport-network/teleport/x/xibc/exported"

	"github.com/teleport-network/teleport-relayer/app/types"
	"github.com/teleport-network/teleport-relayer/app/types/errors"
	"github.com/teleport-network/teleport-relayer/app/utils"
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
		return fmt.Errorf("client height not updated for a long time")
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
	if updateHeight < chainAHeight-1 {
		reqHeight = chainAHeight - 1
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

func (c *Channel) batchGetBlockHeader(reqHeight, revisionHeight, revisionNumber, batchSize uint64) ([]exported.Header, error) {
	times := batchSize
	headers := make([]exported.Header, times)
	var l sync.Mutex
	var wg sync.WaitGroup
	wg.Add(int(times))
	for i := reqHeight; i < reqHeight+times; i++ {
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
	// relay jobs
	relayJobs, err := s.Every(int(c.relayFrequency)).Seconds().Do(func() {
		// when bridge got exception do not to relay packets
		if c.bridgeEnable {
			status, err := utils.RetryGetBridgeStatus(c.bridgeStatusApi)
			if err != nil || status != 1 {
				c.logger.Errorf("Bridge status got exception,should stop relay job, err : %+v", err)
				return
			}
		}
		time.Sleep(time.Duration(c.extraWait*c.relayFrequency) * time.Second)
		c.UpdateHeight()
		if err := c.RelayPackets(c.relayHeight); err != nil {
			c.logger.Errorf("RelayPackets err : %+v", err)
			return
		}
	})
	if err != nil {
		c.logger.Fatal(fmt.Errorf("new relay Jobs error:%+v", err))
	}
	relayJobs.SingletonMode()

	// update client jobs
	if c.chainA.ChainType() == types.ETH || c.chainA.ChainType() == types.BSC {
		updateJobs, err := s.Every(int(c.relayFrequency)).Seconds().Do(func() {
			time.Sleep(time.Duration(c.extraWait*c.relayFrequency) * time.Second)
			if err := c.evmClientUpdate(); err != nil {
				c.logger.Errorf("EvmClientUpdate err : %+v", err)
				return
			}
		})
		if err != nil {
			c.logger.Fatal(fmt.Errorf("new evmClientUpdate Jobs error:%+v", err))
		}
		updateJobs.SingletonMode()
		checkJob, err := s.Every(10).Minute().Do(func() {
			if err := c.checkClient(); err != nil {
				c.logger.Errorf("checkClient err : %+v", err)
			}
		})
		if err != nil {
			c.logger.Fatal("new checkJob error:", err)
		}
		checkJob.SingletonMode()
	}
}

func (c *Channel) RelayPackets(height uint64) error {
	if err := c.handleErrRelayRecord(); err != nil {
		c.logger.Errorf("handleErrRelayRecord error:%+v", err)
	}
	now := time.Now()
	c.logger.Infoln("startRelay ...", now)
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
	var pkts []sdk.Msg
	verifyHeight := c.clientHeight
	if c.chainA.ChainType() == types.Tendermint {
		verifyHeight = chainAHeight
	}
	if height+delayHeight+50 < verifyHeight {
		c.logger.Infof("get packet fromBlock:%v,toBlock:%v", height, height+9)
		pkts, err = c.GetMsg(height, height+49)
		if err != nil {
			return fmt.Errorf("batchGetMsg error:%+v", err)
		}
		updateHeight += 50
	} else if height+delayHeight+c.batchSize < verifyHeight {
		c.logger.Infof("get packet fromBlock:%v,toBlock:%v", height, height+9)
		pkts, err = c.GetMsg(height, height+c.batchSize-1)
		if err != nil {
			return fmt.Errorf("batchGetMsg error:%+v", err)
		}
		updateHeight += c.batchSize
	} else if height+delayHeight < verifyHeight {
		c.logger.Infof("get packet fromBlock:%v,toBlock:%v", height, height)
		pkts, err = c.GetMsg(height, verifyHeight-delayHeight-1)
		if err != nil {
			return fmt.Errorf("get msg err:%+v", err)
		}
		updateHeight += verifyHeight - delayHeight - height
	} else {
		time.Sleep(10 * time.Second)
		return fmt.Errorf("height + delayHeight >= verifyHeight")
	}
	if len(pkts) == 0 {
		c.relayHeight = updateHeight
		return nil
	}
	if c.chainA.ChainType() == types.Tendermint {
		if err := c.UpdateClientByHeight(updateHeight); err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}

	for _, pkt := range pkts {
		// get height , show it when relay success
		chainBHeight, err := c.chainB.GetLatestHeight()
		if err != nil {
			return fmt.Errorf("get chainB latest heigt error:%v", err)
		}

		res, err := c.RetryRelay(pkt)
		if err != nil {
			c.logger.Infof("RelayPackets result: %v , err : %v", res, err)
			packetDetail := types.GetPacketDetail(pkt)
			packetDetail.FromHeight = height
			packetDetail.ToHeight = updateHeight
			packetDetail.ChainName = c.chainA.ChainName()
			packetDetail.ErrMsg = err.Error()
			// Write Err relay details and
			if err := c.errRelay.WriteErrRelay([]types.PacketDetail{packetDetail}, false); err != nil {
				c.logger.Errorf("errRelay.WriteErrRelay error:%v", err.Error())
			}
			continue
		}
		c.logger.Infof(" recv hash : %v ,recv height %v", res, chainBHeight)
	}
	c.logger.Infoln("endRelay ...", time.Now())
	c.relayHeight = updateHeight
	return nil
}

func (c *Channel) handleErrRelayRecord() error {
	var errPacketDetails []types.PacketDetail
	packets, err := c.errRelay.GetErrRelay()
	if err != nil {
		return fmt.Errorf("GetErrRelay err:%+v", err)
	}
	for _, packet := range packets {
		pkts, err := c.GetMsg(packet.FromHeight, packet.ToHeight)
		if err != nil {
			return fmt.Errorf("get msg err:%+v", err)
		}
		for _, pkt := range pkts {
			res, err := c.RetryRelay(pkt)
			if err != nil {
				c.logger.Infof("RelayPackets result: %v , err : %v", res, err)
				packetDetail := types.GetPacketDetail(pkt)
				packetDetail.FromHeight = packet.FromHeight
				packetDetail.ToHeight = packet.ToHeight
				packetDetail.ChainName = c.chainA.ChainName()
				packetDetail.ErrMsg = err.Error()
				errPacketDetails = append(errPacketDetails, packetDetail)
				// Write Err relay details and
				continue
			}
		}
	}
	c.logger.Infoln("errPacketDetails:", errPacketDetails)
	if err := c.errRelay.WriteErrRelay(errPacketDetails, true); err != nil {
		c.logger.Errorf("errRelay.WriteErrRelay error:%v", err.Error())
	}
	return nil
}

func (c *Channel) RetryRelay(pkt sdk.Msg) (res string, err error) {
	for i := 0; i < types.RetryTimes; i++ {
		res, err = c.chainB.RelayPackets(pkt)
		if err != nil {
			// check if packet already relayed
			if handleRecvPacketsError(err) {
				return res, nil
			}
			continue
		}
		break
	}
	return
}

func (c *Channel) ManualRelay(packetRelay *types.PacketDetail, hash string) error {
	start := time.Now()

	latestHeight, err := c.chainA.GetLatestHeight()
	if err != nil {
		return err
	}

	if c.chainA.ChainType() == types.Tendermint {
		c.logger.Infof(" Tendermint client need update client first,update height : %v", latestHeight)
		if err := c.UpdateClientByHeight(latestHeight); err != nil {
			return err
		}
		time.Sleep(3 * time.Second)
	}

	state, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return err
	}

	clientHeight := state.GetLatestHeight().GetRevisionHeight()
	delayHeight := state.GetDelayBlock()
	endHeight := clientHeight - delayHeight

	if packetRelay.ToHeight > endHeight {
		return sdkerrors.Wrapf(errors.ErrDelayHeight, "height must lower than %d ,your hright is %d-%d", endHeight, packetRelay.FromHeight, packetRelay.ToHeight)
	}
	var pkts []sdk.Msg
	if hash != "" {
		pkts, err = c.GetMsgByHash(hash)
		if err != nil {
			return sdkerrors.Wrapf(errors.ErrGetPackets, " err GetMsg by hash! : %v", err)
		}
		_, err = c.manualRelayAll(pkts)
		if err != nil {
			return err
		}
		return nil
	} else {
		pkts, err = c.GetMsg(packetRelay.FromHeight, packetRelay.ToHeight)
		if err != nil {
			return sdkerrors.Wrapf(errors.ErrGetPackets, " err GetMsg! : %v", err)
		}
	}
	if len(pkts) == 0 {
		c.logger.Printf(" no packets need to relay in %v - %v ", packetRelay.FromHeight, packetRelay.ToHeight)
		return nil
	}
	if packetRelay.Sequence != 0 && packetRelay.SrcChain != "" && packetRelay.DestChain != "" {
		err = c.manualRelaySin(pkts, packetRelay.SrcChain, packetRelay.DestChain, packetRelay.RelayChain, packetRelay.Sequence)
		if err != nil {
			return err
		}
	} else {
		_, err = c.manualRelayAll(pkts)
		if err != nil {
			return err
		}
	}
	c.logger.Infof(" relay use time : %v", time.Now().Sub(start))
	return nil
}

func (c *Channel) manualRelaySin(pkts []sdk.Msg, srcChain, destChain, relayChain string, sequence uint64) error {
	receipt, err := c.chainB.GetReceiptPacket(srcChain, destChain, sequence)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrGetPackets, " err get receipt! : %v", err)
	}
	if receipt {
		return sdkerrors.Wrapf(errors.ErrGetPackets, "packet already relay and the receipt is %v ", receipt)
	}
	for _, pkt := range pkts {
		packetDetail := types.GetPacketDetail(pkt)
		if packetDetail.Equal(srcChain, destChain, relayChain, sequence) {
			res, err := c.RetryRelay(pkt)
			if err != nil {
				return sdkerrors.Wrapf(errors.ErrRecvPacket, "packet already relay and the receipt is %v ", receipt)
			}
			c.logger.Infof(" recv hash : %v", res)
			return nil
		}
	}
	return nil
}

func (c *Channel) manualRelayAll(pkts []sdk.Msg) (res string, err error) {
	for _, pkt := range pkts {
		packetDetail := types.GetPacketDetail(pkt)
		receipt, err := c.chainB.GetReceiptPacket(packetDetail.SrcChain, packetDetail.DestChain, packetDetail.Sequence)
		if err != nil {
			return "", sdkerrors.Wrapf(errors.ErrGetReceiptPacket, " err get receipt! : %v", err)
		}
		if receipt {
			c.logger.Infof(" packet already relay the receipt is %v \n packet detiles:", receipt)
			continue
		}
		res, err = c.RetryRelay(pkt)
		if err != nil {
			return "", sdkerrors.Wrapf(errors.ErrRecvPacket, " err : %v", err)
		}
		c.logger.Infof(" recv hash : %v", res)
	}
	return
}
