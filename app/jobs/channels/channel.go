package channels

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"

	interfaces2 "github.com/teleport-network/teleport-relayer/app/interfaces"

	log "github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport/x/xibc/exported"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/repo/cache"
	"github.com/teleport-network/teleport-relayer/app/types"
)

var _ IChannel = new(Channel)

type IChannel interface {
	RelayTask(s *gocron.Scheduler)
	EvmClientUpdate() error
}

type Channel struct {
	chainA       interfaces2.IChain
	chainB       interfaces2.IChain
	relayHeight  uint64
	clientHeight uint64 // TODO
	chainName    string
	state        *cache.CacheFileWriter
	logger       *log.Logger
}

// TODO: pullHeight   relayHeight   chainAHeight   clientAHeight

func NewChannel(
	source interfaces2.IChain,
	dest interfaces2.IChain,
	height uint64,
	cacheName string,
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
		chainA:      source,
		chainB:      dest,
		relayHeight: startHeight,
		chainName:   source.ChainName(),
		state:       state,
		logger:      logger,
	}, nil
}

func (c *Channel) UpdateHeight() {
	if err := c.state.Write(c.relayHeight);err != nil {
		panic(err)
	}
}

func (c *Channel) EvmClientUpdate() error {
	if c.chainA.ChainType() == types.Tendermint {
		return nil
	}
	chainAHeight, _ := c.chainA.GetLatestHeight()
	clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
	if err != nil {
		return err
	}
	updateHeight := clientState.GetLatestHeight().GetRevisionHeight() + 1
	if updateHeight >= chainAHeight {
		return fmt.Errorf("updateHeight %v > chainA height %v,no use update", updateHeight, chainAHeight)
	}
	c.logger.Println("update client updateHeight:", updateHeight)
	for {
		var header exported.Header
		req := &types.GetBlockHeaderReq{
			LatestHeight:   updateHeight,
			TrustedHeight:  clientState.GetLatestHeight().GetRevisionHeight(),
			RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
		}
		t1 := time.Now()
		header, err = c.chainA.GetBlockHeader(req)
		if err != nil {
			c.logger.Println("get blockHeader err", err)
			continue
		}
		t2 := time.Now()
		duration := t2.Sub(t1)
		c.logger.Println("get BlockHeader duration:", duration)
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
