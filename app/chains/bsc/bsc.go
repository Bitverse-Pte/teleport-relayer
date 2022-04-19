package bsc

import (
	log "github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport-relayer/app/config"
	interfaces2 "github.com/teleport-network/teleport-relayer/app/interfaces"
	"github.com/teleport-network/teleport-relayer/app/types"
)

func InitBscChain(cfg *config.ChainCfg, logger *log.Logger) interfaces2.IChain {
	loggerEntry := logger.WithFields(log.Fields{
		"chain_name": cfg.Tendermint.ChainName,
	})

	loggerEntry.Info(" init bsc chain start")

	contractCfgGroup := NewContractCfgGroup()
	contractCfgGroup.Packet.Addr = cfg.Bsc.Contracts.Packet.Addr
	contractCfgGroup.Packet.Topic = cfg.Bsc.Contracts.Packet.Topic
	contractCfgGroup.Packet.OptPrivKey = cfg.Bsc.Contracts.Packet.OptPrivKey

	contractCfgGroup.AckPacket.Addr = cfg.Bsc.Contracts.AckPacket.Addr
	contractCfgGroup.AckPacket.Topic = cfg.Bsc.Contracts.AckPacket.Topic

	contractCfgGroup.Client.Addr = cfg.Bsc.Contracts.Client.Addr
	contractCfgGroup.Client.Topic = cfg.Bsc.Contracts.Client.Topic
	contractCfgGroup.Client.OptPrivKey = cfg.Bsc.Contracts.Client.OptPrivKey

	contractBindOptsCfg := NewContractBindOptsCfg()
	contractBindOptsCfg.ChainID = cfg.Bsc.ChainID
	contractBindOptsCfg.ClientPrivKey = cfg.Bsc.Contracts.Client.OptPrivKey
	contractBindOptsCfg.PacketPrivKey = cfg.Bsc.Contracts.Packet.OptPrivKey
	contractBindOptsCfg.GasLimit = cfg.Bsc.GasLimit
	contractBindOptsCfg.MaxGasPrice = cfg.Bsc.MaxGasPrice

	bscChainCfg := NewChainConfig()
	bscChainCfg.ContractCfgGroup = contractCfgGroup
	bscChainCfg.ContractBindOptsCfg = contractBindOptsCfg

	bscChainCfg.ChainType = types.BSC
	bscChainCfg.ChainName = cfg.Bsc.ChainName
	bscChainCfg.ChainID = cfg.Bsc.ChainID
	bscChainCfg.ChainURI = cfg.Bsc.URI
	bscChainCfg.Slot = cfg.Bsc.CommentSlot
	bscChainCfg.UpdateClientFrequency = cfg.Bsc.UpdateClientFrequency
	bscChainCfg.TipCoefficient = cfg.Bsc.TipCoefficient

	bscRepo, err := NewBsc(bscChainCfg)
	if err != nil {
		logger.WithFields(
			log.Fields{
				"chain_name": cfg.Tendermint.ChainName,
				"err_msg":    err,
			},
		).Fatal("failed to init chain")
	}

	return bscRepo
}
