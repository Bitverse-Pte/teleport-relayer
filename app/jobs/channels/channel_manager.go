package channels

import (
	"github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport-relayer/app/chains/bsc"
	"github.com/teleport-network/teleport-relayer/app/chains/eth"
	"github.com/teleport-network/teleport-relayer/app/chains/tendermint"
	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/interfaces"
)

const TendermintAndTendermint = "tendermint_and_tendermint"
const TendermintAndETH = "tendermint_and_eth"
const TendermintAndBsc = "tendermint_and_bsc"

func NewChannelMap(cfg *config.Config, logger *logrus.Logger) map[string]IChannel {
	if len(cfg.App.ChannelTypes) != 1 {
		logger.Fatal("channel_types should be equal 1")
	}
	for _, channelType := range cfg.App.ChannelTypes {
		switch channelType {
		case TendermintAndTendermint:
			sourceChain := tendermint.InitTendermintChain(&cfg.Chain.Source, logger)
			destChain := tendermint.InitTendermintChain(&cfg.Chain.Dest, logger)
			return MakeChannels(cfg, sourceChain, destChain, logger)
		case TendermintAndETH:
			sourceChain := tendermint.InitTendermintChain(&cfg.Chain.Source, logger)
			destChain := eth.InitEthChain(&cfg.Chain.Dest, logger)
			return MakeChannels(cfg, sourceChain, destChain, logger)
		case TendermintAndBsc:
			sourceChain := tendermint.InitTendermintChain(&cfg.Chain.Source, logger)
			destChain := bsc.InitBscChain(&cfg.Chain.Dest, logger)
			return MakeChannels(cfg, sourceChain, destChain, logger)
		default:
			logger.WithFields(logrus.Fields{"channel_type": channelType}).Fatal("CreateChannel type does not exist")
		}
	}
	return nil
}

func MakeChannels(cfg *config.Config, sourceChain, destChain interfaces.IChain, logger *logrus.Logger) map[string]IChannel {
	srcChannel, err := NewChannel(sourceChain, destChain, cfg.Chain.Source.Cache.StartHeight, cfg.Chain.Source.Cache.Filename, logger)
	if err != nil {
		panic(err)
	}
	destChannel, err := NewChannel(destChain, sourceChain, cfg.Chain.Dest.Cache.StartHeight, cfg.Chain.Dest.Cache.Filename, logger)
	if err != nil {
		panic(err)
	}
	channelMap := map[string]IChannel{
		sourceChain.ChainName(): srcChannel,
		destChain.ChainName():   destChannel,
	}
	return channelMap
}
