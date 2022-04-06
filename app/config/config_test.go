package config

import (
	"fmt"
	"testing"

	"github.com/teleport-network/teleport-relayer/tools"
)

func TestLoadConfigs(t *testing.T) {
	cfg := Config{}
	tools.InitTomlConfigs([]*tools.ConfigMap{
		{
			FilePath: "ethconfig.toml",
			Pointer:  &cfg,
		},
	})
	fmt.Println(cfg.Chain.Source.Tendermint.GrpcAddr1)
}
