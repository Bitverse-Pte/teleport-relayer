package channels

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-co-op/gocron"

	log "github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport/x/xibc/exported"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/dto"
	"github.com/teleport-network/teleport-relayer/app/interfaces"
	"github.com/teleport-network/teleport-relayer/app/repo/cache"
	"github.com/teleport-network/teleport-relayer/app/types"
)

var _ IChannel = new(Channel)

const DefaultBatchSize uint64 = 10

type IChannel interface {
	RelayTask(s *gocron.Scheduler)
	EvmClientUpdate(s *gocron.Scheduler)
	UpgradeRelayHeight(ctx *gin.Context)
	ViewRelayHeight(ctx *gin.Context)
	UpgradeExtraWait(ctx *gin.Context)
	ViewExtraWait(ctx *gin.Context)
	ManualRelayByHash(ctx *gin.Context)
	ManualRelay(detail *types.PacketDetail, hash string) error
}

type Channel struct {
	chainA          interfaces.IChain
	chainB          interfaces.IChain
	relayHeight     uint64
	clientHeight    uint64
	checkHeight     uint64
	chainName       string
	relayFrequency  uint64
	extraWait       uint64 // waitTime = (extraWait * relayFrequency) second
	state           *cache.CacheFileWriter
	errRelay        *cache.CacheFileWriter
	logger          *log.Logger
	batchSize       uint64
	bridgeStatusApi string
	bridgeEnable    bool
}

// TODO: pullHeight   relayHeight   chainAHeight   clientAHeight

func NewChannel(
	chainA interfaces.IChain,
	chainB interfaces.IChain,
	chainACfg config.ChainCfg,
	logger *log.Logger,
	cfg *config.Config,
) (
	IChannel,
	error,
) {
	var startHeight uint64
	state := cache.NewCacheFileWriter(config.Home, config.DefaultCacheDirName, chainACfg.Cache.Filename)
	errRelayCache := cache.NewCacheFileWriter(config.Home, config.DefaultCacheDirName, "errRelay")
	stateData := state.LoadCache()
	if chainA.ChainType() == types.Tendermint {
		startHeight = chainACfg.Cache.StartHeight
	} else {
		clientStatus, err := chainB.GetLightClientState(chainA.ChainName())
		if err != nil {
			return nil, err
		}
		startHeight = clientStatus.GetLatestHeight().GetRevisionHeight() + 1
	}
	if stateData.LatestHeight != 0 {
		startHeight = stateData.LatestHeight
	}
	if chainACfg.Cache.RevisedHeight != 0 {
		if err := state.Write(chainACfg.Cache.RevisedHeight); err != nil {
			panic(fmt.Errorf("state.Write revisedHeight error:%+v", err))
		}
		startHeight = chainACfg.Cache.RevisedHeight
	}
	batchSize := DefaultBatchSize
	if chainACfg.BatchSize != 0 {
		batchSize = chainACfg.BatchSize
	}
	return &Channel{
		chainA:          chainA,
		chainB:          chainB,
		relayHeight:     startHeight,
		chainName:       chainA.ChainName(),
		relayFrequency:  chainACfg.RelayFrequency,
		state:           state,
		logger:          logger,
		batchSize:       batchSize,
		bridgeStatusApi: cfg.App.BridgeStatusApi,
		bridgeEnable:    cfg.App.BridgeEnable,
		errRelay:        errRelayCache,
	}, nil
}

func (c *Channel) UpdateHeight() {
	if err := c.state.Write(c.relayHeight); err != nil {
		panic(fmt.Errorf("state.Write error:%+v", err))
	}
}

func (c *Channel) WriteErrRelay() {
	if err := c.state.Write(c.relayHeight); err != nil {
		panic(fmt.Errorf("WriteErrRelay.Write error:%+v", err))
	}
}

func (c *Channel) UpgradeRelayHeight(ctx *gin.Context) {
	var heightObj dto.ReqUpgradeHeight
	if err := ctx.Bind(&heightObj); err != nil {
		ctx.JSON(http.StatusOK, dto.Response{Code: dto.BadRequest, Message: fmt.Sprintf("invalid type:%v", err.Error())})
		return
	}
	if heightObj.Height == 0 {
		ctx.JSON(http.StatusOK, dto.Response{Code: dto.BadRequest, Message: fmt.Sprintf("height = 0")})
		return
	}
	chainAHeight, err := c.chainA.GetLatestHeight()
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{Code: dto.BadRequest, Message: fmt.Sprintf("get latest height err:%v", err.Error())})
		return
	}
	if chainAHeight < heightObj.Height {
		ctx.JSON(http.StatusOK, dto.Response{Code: dto.BadRequest, Message: fmt.Sprintf("upgrade height %v > latest height %v\n", heightObj.Height, chainAHeight)})
		return
	}
	c.relayHeight = heightObj.Height
	ctx.JSON(http.StatusOK, dto.Response{Code: dto.Success, Message: "success", Data: c.relayHeight})
}

func (c *Channel) ViewRelayHeight(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, dto.Response{Code: dto.Success, Message: "success", Data: c.relayHeight})
}

func (c *Channel) ViewExtraWait(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, dto.Response{Code: dto.Success, Message: "success", Data: c.extraWait})
}

func (c *Channel) UpgradeExtraWait(ctx *gin.Context) {
	var extraWaitObj dto.ReqUpgradeExtraWait
	if err := ctx.Bind(&extraWaitObj); err != nil {
		ctx.JSON(http.StatusOK, dto.Response{Code: dto.BadRequest, Message: fmt.Sprintf("invalid type:%v", err.Error())})
		return
	}
	c.extraWait = extraWaitObj.ExtraWait
	ctx.JSON(http.StatusOK, dto.Response{Code: dto.Success, Message: "success", Data: c.extraWait})
}

func (c *Channel) ManualRelayByHash(ctx *gin.Context) {
	var tx dto.ReqRelayByHash
	if err := ctx.Bind(&tx); err != nil {
		ctx.JSON(http.StatusOK, dto.Response{Code: dto.BadRequest, Message: fmt.Sprintf("invalid type:%v", err.Error())})
		return
	}
	pkt, err := c.GetMsgByHash(tx.Hash)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{Code: dto.BadRequest, Message: fmt.Sprintf("get msg error :%v", err.Error())})
		return
	}
	res, err := c.manualRelayAll(pkt)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{Code: dto.BadRequest, Message: fmt.Sprintf("manual relay error :%v", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{Code: dto.Success, Message: "success", Data: fmt.Sprintf("recv hash :%v", res)})
}

func (c *Channel) EvmClientUpdate(s *gocron.Scheduler) {
	if c.chainA.ChainType() == types.ETH || c.chainA.ChainType() == types.BSC {
		jobs, err := s.Every(int(c.relayFrequency)).Seconds().Do(func() {
			time.Sleep(time.Duration(c.extraWait*c.relayFrequency) * time.Second)
			if err := c.evmClientUpdate(); err != nil {
				c.logger.Errorf("EvmClientUpdate err : %+v", err)
				time.Sleep(10 * time.Second)
				return
			}
		})
		if err != nil {
			panic(fmt.Errorf("new EvmClientUpdate jobs error:%+v", err))
		}
		jobs.SingletonMode()
	}
}

func (c *Channel) evmClientUpdate() error {
	chainAHeight, _ := c.chainA.GetLatestHeight()
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return fmt.Errorf("GetLightClientState error:%+v,chainA:%v,chainB:%v", err, c.chainA.ChainName(), c.chainB.ChainName())
	}
	updateHeight := clientState.GetLatestHeight().GetRevisionHeight() + 1
	if updateHeight >= chainAHeight {
		c.logger.Infof("updateHeight %v >= chainA height %v,no use update", updateHeight, chainAHeight)
		time.Sleep(40 * time.Second)
		return nil
	}
	revisionHeight := clientState.GetLatestHeight().GetRevisionHeight()
	revisionNumber := clientState.GetLatestHeight().GetRevisionNumber()
	delayHeight := clientState.GetDelayBlock()
	c.logger.Println("chainAHeight", chainAHeight)
	c.logger.Println("update client updateHeight:", updateHeight)
	if chainAHeight > updateHeight+delayHeight+50 {
		headers, err := c.batchGetBlockHeader(updateHeight, revisionHeight, revisionNumber, 50)
		if err != nil {
			return fmt.Errorf("batchGetBlockHeader error:%+v", err)
		}
		return c.chainB.BatchUpdateClient(headers, c.chainA.ChainName())
	}
	if chainAHeight > updateHeight+delayHeight+c.batchSize {
		headers, err := c.batchGetBlockHeader(updateHeight, revisionHeight, revisionNumber, c.batchSize)
		if err != nil {
			return fmt.Errorf("batchGetBlockHeader error:%+v", err)
		}
		return c.chainB.BatchUpdateClient(headers, c.chainA.ChainName())
	} else if (chainAHeight < updateHeight+delayHeight+c.batchSize) && (chainAHeight > updateHeight+delayHeight) {
		headers, err := c.batchGetBlockHeader(updateHeight, revisionHeight, revisionNumber, chainAHeight-updateHeight-delayHeight)
		if err != nil {
			return fmt.Errorf("batchGetBlockHeader error:%+v", err)
		}
		return c.chainB.BatchUpdateClient(headers, c.chainA.ChainName())
	}
	for {
		var header exported.Header
		req := &types.GetBlockHeaderReq{
			LatestHeight:   updateHeight,
			TrustedHeight:  clientState.GetLatestHeight().GetRevisionHeight(),
			RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
		}
		header, err = c.chainA.GetBlockHeader(req)
		if err != nil {
			c.logger.Println("get blockHeader err", err)
			return fmt.Errorf("get blockHeader err:%+v", err)
		}
		if err = c.chainB.UpdateClient(header, c.chainA.ChainName()); err != nil {
			if isBifurcate(err) {
				updateHeight--
				continue
			}
			return fmt.Errorf("update client err:%+v", err)
		}
		break
	}
	return nil
}
