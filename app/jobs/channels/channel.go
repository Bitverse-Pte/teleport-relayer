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

type IChannel interface {
	RelayTask(s *gocron.Scheduler)
	EvmClientUpdate(s *gocron.Scheduler)
	UpgradeRelayHeight(ctx *gin.Context)
	ViewRelayHeight(ctx *gin.Context)
	UpgradeExtraWait(ctx *gin.Context)
	ViewExtraWait(ctx *gin.Context)
}

type Channel struct {
	chainA         interfaces.IChain
	chainB         interfaces.IChain
	relayHeight    uint64
	clientHeight   uint64 // TODO
	chainName      string
	relayFrequency uint64
	extraWait      uint64  // waitTime = (extraWait * relayFrequency) second
	state          *cache.CacheFileWriter
	logger         *log.Logger
}

// TODO: pullHeight   relayHeight   chainAHeight   clientAHeight

func NewChannel(
	source interfaces.IChain,
	dest interfaces.IChain,
	height uint64,
	cacheName string,
	relayFrequency uint64,
	logger *log.Logger,
) (
	IChannel,
	error,
) {
	var startHeight uint64
	state := cache.NewCacheFileWriter(config.Home, config.DefaultCacheDirName, cacheName)
	stateData := state.LoadCache()
	if source.ChainType() == types.Tendermint {
		startHeight = height
	} else {
		clientStatus, err := dest.GetLightClientState(source.ChainName())
		if err != nil {
			return nil, err
		}
		startHeight = clientStatus.GetLatestHeight().GetRevisionHeight() + 1
	}
	if stateData.LatestHeight != 0 {
		startHeight = stateData.LatestHeight
	}
	return &Channel{
		chainA:         source,
		chainB:         dest,
		relayHeight:    startHeight,
		chainName:      source.ChainName(),
		relayFrequency: relayFrequency,
		state:          state,
		logger:         logger,
	}, nil
}

func (c *Channel) UpdateHeight() {
	if err := c.state.Write(c.relayHeight); err != nil {
		panic(err)
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

func (c *Channel) EvmClientUpdate(s *gocron.Scheduler) {
	if c.chainA.ChainType() == types.ETH || c.chainA.ChainType() == types.BSC {
		s.Every(5).Seconds().Do(func() {
			if err := c.evmClientUpdate(); err != nil {
				c.logger.Errorf("EvmClientUpdate err : %+v", err)
				time.Sleep(10 * time.Second)
				return
			}
		})
	}
}

func (c *Channel) evmClientUpdate() error {
	chainAHeight, _ := c.chainA.GetLatestHeight()
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return err
	}
	updateHeight := clientState.GetLatestHeight().GetRevisionHeight() + 1
	if updateHeight >= chainAHeight {
		c.logger.Infof("updateHeight %v > chainA height %v,no use update", updateHeight, chainAHeight)
		time.Sleep(40 * time.Second)
		return nil
	}
	c.logger.Println("update client updateHeight:", updateHeight)
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
			continue
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
