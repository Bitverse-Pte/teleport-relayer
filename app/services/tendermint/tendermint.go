package tendermint

import (
	log "github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/services/interfaces"
	"github.com/teleport-network/teleport-relayer/app/types"
)

func InitTendermintChain(cfg *config.ChainCfg, logger *log.Logger) interfaces.IChain {
	chainClient, err := NewTendermintClient(
		types.Tendermint,
		cfg.Tendermint.ChainName,
		cfg.Tendermint.UpdateClientFrequency,
		&cfg.Tendermint,
	)
	if err != nil {
		panic(err)
	}
	return chainClient
}
