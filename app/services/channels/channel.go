package channels

import (
	"time"

	log "github.com/sirupsen/logrus"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/teleport-network/teleport/x/xibc/exported"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/repostitory/cache"
	"github.com/teleport-network/teleport-relayer/app/services/interfaces"
	"github.com/teleport-network/teleport-relayer/app/types"
)

var _ IChannel = new(Channel)

const ErrSleep = 3

type IChannel interface {
	Relay()
	EthClientSync()
	PacketSync()
	DeleteRelayedPacket()
}

type Channel struct {
	chainA       interfaces.IChain
	chainB       interfaces.IChain
	relayHeight  uint64
	clientHeight uint64 // TODO
	chainName    string
	state        *cache.CacheFileWriter
	logger       *log.Logger
	PacketPool   *PacketPool
	PacketDBPool *PacketDBPool
}

// TODO: pullHeight   relayHeight   chainAHeight   clientAHeight

func NewChannel(
	source interfaces.IChain,
	dest interfaces.IChain,
	height uint64,
	cacheName string,
	logger *log.Logger,
) (
	IChannel,
	error,
) {
	var startHeight uint64
	state := cache.NewCacheFileWriter(config.DefaultHomePath, config.DefaultCacheDirName, cacheName)
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
	bp := make(map[uint64][]sdk.Msg)
	packetPool := &PacketPool{
		syncHeight: startHeight, // TODO
		BP:         bp,
	}
	return &Channel{
		chainA:      source,
		chainB:      dest,
		relayHeight: startHeight,
		chainName:   source.ChainName(),
		state:       state,
		PacketPool:  packetPool,
		logger:      logger,
	}, nil
}

func (c *Channel) UpdateHeight() {
	_ = c.state.Write(c.relayHeight)
}

func (c *Channel) Relay() {
	for {
		c.UpdateHeight()
		if err := c.relay(); err != nil {
			time.Sleep(ErrSleep * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func (c *Channel) EthClientSync() {
	defer func() {
		if err := recover(); err != nil {
			c.logger.Printf("EthClientSync  panic:%v", err)
		}
	}()
	if c.chainA.ChainType() != types.ETH {
		return
	}
	var updateHeight uint64
	for {
		chainAHeight, _ := c.chainA.GetLatestHeight()
		if updateHeight >= chainAHeight {
			time.Sleep(3 * ErrSleep * time.Second)
			continue
		}
		clientState, err := c.chainB.GetLightClientState(c.chainA.ChainName())
		if err != nil {
			continue
		}
		updateHeight = clientState.GetLatestHeight().GetRevisionHeight() + 1
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
				c.logger.Printf("update client err:%+v", err)
			}
			break
		}
	}
}

func (c *Channel) PacketSync() {
	for {
		if err := c.packetSync(); err != nil {
			c.logger.Printf("packetSync ERROR:%+v", err)
			time.Sleep(ErrSleep * time.Second)
		}
	}
}

func (c *Channel) PacketSyncToDB() {
	for {
		if err := c.packetSyncToDB(); err != nil {
			time.Sleep(ErrSleep * time.Second)
		}
	}
}
