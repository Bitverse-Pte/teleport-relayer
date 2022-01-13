package eth

import (
	log "github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/services/interfaces"
	"github.com/teleport-network/teleport-relayer/app/types"
)

func InitEthChain(cfg *config.ChainCfg, logger *log.Logger) interfaces.IChain {
	loggerEntry := logger.WithFields(log.Fields{
		"chain_name": cfg.Tendermint.ChainName,
	})

	loggerEntry.Info(" init eth chain start")

	contractCfgGroup := NewContractCfgGroup()
	contractCfgGroup.Packet.Addr = cfg.Eth.Contracts.Packet.Addr
	contractCfgGroup.Packet.Topic = cfg.Eth.Contracts.Packet.Topic
	contractCfgGroup.Packet.OptPrivKey = cfg.Eth.Contracts.Packet.OptPrivKey

	contractCfgGroup.AckPacket.Addr = cfg.Eth.Contracts.AckPacket.Addr
	contractCfgGroup.AckPacket.Topic = cfg.Eth.Contracts.AckPacket.Topic

	contractCfgGroup.Client.Addr = cfg.Eth.Contracts.Client.Addr
	contractCfgGroup.Client.Topic = cfg.Eth.Contracts.Client.Topic
	contractCfgGroup.Client.OptPrivKey = cfg.Eth.Contracts.Client.OptPrivKey

	contractBindOptsCfg := NewContractBindOptsCfg()
	contractBindOptsCfg.ChainID = cfg.Eth.ChainID
	contractBindOptsCfg.ClientPrivKey = cfg.Eth.Contracts.Client.OptPrivKey
	contractBindOptsCfg.PacketPrivKey = cfg.Eth.Contracts.Packet.OptPrivKey
	contractBindOptsCfg.GasLimit = cfg.Eth.GasLimit
	contractBindOptsCfg.MaxGasPrice = cfg.Eth.MaxGasPrice

	ethChainCfg := NewChainConfig()
	ethChainCfg.ContractCfgGroup = contractCfgGroup
	ethChainCfg.ContractBindOptsCfg = contractBindOptsCfg

	ethChainCfg.ChainType = types.ETH
	ethChainCfg.ChainName = cfg.Eth.ChainName
	ethChainCfg.ChainID = cfg.Eth.ChainID
	ethChainCfg.ChainURI = cfg.Eth.URI
	ethChainCfg.Slot = cfg.Eth.CommentSlot
	ethChainCfg.UpdateClientFrequency = cfg.Eth.UpdateClientFrequency
	ethChainCfg.TipCoefficient = cfg.Eth.TipCoefficient

	ethRepo, err := NewEth(ethChainCfg)
	if err != nil {
		logger.WithFields(
			log.Fields{
				"chain_name": cfg.Tendermint.ChainName,
				"err_msg":    err,
			},
		).Fatal("failed to init chain")
	}

	return ethRepo
}
