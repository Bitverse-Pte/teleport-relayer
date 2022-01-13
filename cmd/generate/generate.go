package generate

import (
	"bufio"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/services/tendermint"
)

const (
	TendermintAndTendermint    = "tendermint_and_tendermint"
	TendermintAndETH           = "tendermint_and_eth"
	xibcTendermintMerklePrefix = "xibc"
	xibcTendermintRoot         = "app_hash"
	ETH                        = "eth"
	TENDERMINT                 = "tendermint"
)

const (
	clientStatePrefix       = `{"@type":"/xibc.clients.lightclients.tendermint.v1.ClientState",`
	consensusStatePrefix    = `{"@type":"/xibc.clients.lightclients.tendermint.v1.ConsensusState",`
	EthConsensusStatePrefix = `{"@type":"/xibc.clients.lightclients.eth.v1.ConsensusState",`
	EthClientStatePrefix    = `{"@type":"/xibc.clients.lightclients.eth.v1.ClientState",`
)

func GenerateClientFiles() {
	cfg := config.LoadConfigs()
	for _, channelType := range cfg.App.ChannelTypes {
		switch channelType {
		case TendermintAndTendermint:
			logger := logrus.WithFields(logrus.Fields{
				"source_chain": &cfg.Chain.Source.Tendermint.ChainName,
				"dest_chain":   &cfg.Chain.Dest.Tendermint.ChainName,
			})

			logger.Info("1. init source chain")
			// sourceChain := tendermintCreateClientFiles(&cfg.Chain.Source, logger)
			// sourceChain := tendermint.InitTendermintChain(&cfg.Chain.Source,nil)
			sourceChain, err := tendermint.NewTendermintClient(
				cfg.Chain.Source.ChainType,
				cfg.Chain.Source.Tendermint.ChainName,
				cfg.Chain.Source.Tendermint.UpdateClientFrequency,
				&cfg.Chain.Source.Tendermint,
			)
			if err != nil {
				return
			}
			generateTendermintJson(
				sourceChain,
				int64(cfg.Chain.Source.Cache.StartHeight),
				cfg.Chain.Source.Tendermint.ChainName,
			)
			logger.Info("2. init dest chain")
			// destChain := tendermintCreateClientFiles(&cfg.Chain.Dest, logger)
			// destChain := tendermint.InitTendermintChain(&cfg.Chain.Dest,nil)
			destChain, _ := tendermint.NewTendermintClient(
				cfg.Chain.Dest.ChainType,
				cfg.Chain.Dest.Tendermint.ChainName,
				cfg.Chain.Dest.Tendermint.UpdateClientFrequency,
				&cfg.Chain.Dest.Tendermint,
			)
			generateTendermintJson(
				destChain,
				int64(cfg.Chain.Dest.Cache.StartHeight),
				cfg.Chain.Dest.Tendermint.ChainName,
			)
		case TendermintAndETH:

			if cfg.Chain.Source.ChainType == TENDERMINT && cfg.Chain.Dest.ChainType == ETH {
				logger := logrus.WithFields(logrus.Fields{
					"source_chain": &cfg.Chain.Source.Tendermint.ChainName,
					"dest_chain":   &cfg.Chain.Dest.Eth.ChainName,
				})
				logger.Info("1. init source chain")
				// sourceChain := tendermintCreateClientFiles(&cfg.Chain.Source, logger)
				sourceChain, err := tendermint.NewTendermintClient(
					cfg.Chain.Source.ChainType,
					cfg.Chain.Source.Tendermint.ChainName,
					cfg.Chain.Source.Tendermint.UpdateClientFrequency,
					&cfg.Chain.Source.Tendermint,
				)
				if err != nil {
					// TODO
					return
				}
				generateTendermintHex(
					sourceChain,
					int64(cfg.Chain.Source.Cache.StartHeight),
					cfg.Chain.Source.Tendermint.ChainName,
					logger,
				)
				logger.Info("2. init dest chain")
				generateETHJson(&cfg.Chain.Dest, sourceChain, logger)
			}

			if cfg.Chain.Source.ChainType == ETH && cfg.Chain.Dest.ChainType == TENDERMINT {
				logger := logrus.WithFields(logrus.Fields{
					"source_chain": &cfg.Chain.Source.Eth.ChainName,
					"dest_chain":   &cfg.Chain.Dest.Tendermint.ChainName,
				})
				logger.Info("1. init dest  chain")
				// destChain := tendermintCreateClientFiles(&cfg.Chain.Dest, logger)
				// destChain := tendermint.InitTendermintChain(&cfg.Chain.Dest, logger)
				destChain, err := tendermint.NewTendermintClient(
					cfg.Chain.Dest.ChainType,
					cfg.Chain.Dest.Tendermint.ChainName,
					cfg.Chain.Dest.Tendermint.UpdateClientFrequency,
					&cfg.Chain.Dest.Tendermint,
				)
				if err != nil {
					// TODO
					return
				}
				generateTendermintHex(
					destChain,
					int64(cfg.Chain.Dest.Cache.StartHeight),
					cfg.Chain.Dest.Tendermint.ChainName,
					logger,
				)
				logger.Info("2. init source chain")
				generateETHJson(&cfg.Chain.Source, destChain, logger)
			}

		}
	}
}

func WriteCreateClientFiles(fileName string, content string) {
	home := config.Home
	if home == "" {
		home = config.DefaultHomePath
	}
	path := filepath.Join(home, fileName)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(content); err != nil {
		panic(err)
	}
	writer.Flush()
}

// exponential backoff (with jitter)
// 0.5s -> 2s -> 4.5s -> 8s -> 12.5 with 1s variation
func backoffTimeout(attempt uint16) time.Duration {
	// nolint:gosec // G404: Use of weak random number generator
	return time.Duration(500*attempt*attempt)*time.Millisecond + time.Duration(rand.Intn(1000))*time.Millisecond
}
