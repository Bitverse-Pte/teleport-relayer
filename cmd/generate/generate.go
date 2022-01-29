package generate

import (
	"bufio"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/teleport-network/teleport-sdk-go/client"

	"github.com/teleport-network/teleport-relayer/app/chains/tendermint"

	"github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport-relayer/app/config"
)

const (
	TendermintAndTendermint = "tendermint_and_tendermint"
	TendermintAndETH        = "tendermint_and_eth"
	TendermintAndBsc        = "tendermint_and_bsc"

	xibcTendermintMerklePrefix = "xibc"
	xibcTendermintRoot         = "app_hash"

	ETH        = "eth"
	TENDERMINT = "tendermint"
	Bsc        = "bsc"
)

var (
	OutPut     string
	Type       string
	Height     int64
	GrpcAddr   string
	ChainID    string
	PacketAddr string
	ChainName  string
)

func GenClientFiles() {
	if Type == "" || ChainName == "" || GrpcAddr == "" {
		return
	}
	config.Home = OutPut
	if OutPut == "" {
		config.Home = config.DefaultHomePath
	}
	switch Type {
	case TENDERMINT:
		logger := logrus.WithFields(logrus.Fields{
			"type": TENDERMINT,
		})
		if Height == 0 {
			logger.Error("tendermint height 0")
			return
		}
		sdkCli, err := client.NewClient(GrpcAddr, ChainID)
		if err != nil {
			panic(err)
		}
		cli := &tendermint.Tendermint{
			Codec:       tendermint.MakeCodec(),
			TeleportSDK: sdkCli,
		}
		generateTendermintJson(cli, Height, Type, logger)
		generateTendermintHex(cli, Height, Type, logger)

	case ETH:
		logger := logrus.WithFields(logrus.Fields{
			"type": ETH,
		})
		if PacketAddr == "" {
			logger.Error("packet address empty")
			return
		}
		chainId, err := strconv.ParseInt(ChainID, 10, 64)
		if err != nil {
			logger.Error("chainId must be unit64")
			return
		}
		eth := &config.Eth{
			URI:       GrpcAddr,
			ChainID:   uint64(chainId),
			ChainName: ChainName,
			Contracts: config.EthContracts{
				Packet: config.EthContractCfg{
					Addr: PacketAddr,
				},
			},
		}
		generateETHJson(eth, tendermint.MakeCodec(), logger)

	case Bsc:
		logger := logrus.WithFields(logrus.Fields{
			"type": Bsc,
		})
		if PacketAddr == "" {
			logger.Error("packet address empty")
			return
		}
		chainId, err := strconv.ParseInt(ChainID, 10, 64)
		if err != nil {
			logger.Error("chainId must be unit64")
			return
		}
		bsc := &config.Bsc{
			URI:       GrpcAddr,
			ChainID:   uint64(chainId),
			ChainName: ChainName,
			Contracts: config.EthContracts{
				Packet: config.EthContractCfg{
					Addr: PacketAddr,
				},
			},
		}
		generateBscJson(bsc, tendermint.MakeCodec(), logger)

	}
}

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
				logger,
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
				logger,
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
				generateETHJson(&cfg.Chain.Dest.Eth, sourceChain.Codec, logger)
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
				generateETHJson(&cfg.Chain.Source.Eth, destChain.Codec, logger)
			}

		case TendermintAndBsc:
			if cfg.Chain.Source.ChainType == TENDERMINT && cfg.Chain.Dest.ChainType == Bsc {
				logger := logrus.WithFields(logrus.Fields{
					"source_chain": &cfg.Chain.Source.Tendermint.ChainName,
					"dest_chain":   &cfg.Chain.Dest.Bsc.ChainName,
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
				generateBscJson(&cfg.Chain.Dest.Bsc, sourceChain.Codec, logger)
			}

			if cfg.Chain.Source.ChainType == Bsc && cfg.Chain.Dest.ChainType == TENDERMINT {
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
				generateBscJson(&cfg.Chain.Source.Bsc, destChain.Codec, logger)
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
