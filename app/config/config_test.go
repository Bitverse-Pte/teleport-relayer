package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

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
	marshal, err := json.Marshal(cfg)
	require.NoError(t, err)
	t.Log(string(marshal))
}
