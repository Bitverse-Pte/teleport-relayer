package tendermint

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/interfaces"
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
		panic(fmt.Errorf("NewTendermintClient error:%+v", err))
	}
	return chainClient
}
