package config

import (
	"os"
	"path"
	"path/filepath"

	"github.com/teleport-network/teleport-relayer/tools"
)

const (
	DefaultHomeDirName   = ".teleport-relayer"
	DefaultConfigDirName = "configs" // TODO delete initialization.DefaultConfigDirName
	DefaultConfigName    = "config.toml"
	DefaultCacheDirName  = "cache"
)

var (
	Home            string
	LocalConfig     string
	UserDir, _      = os.UserHomeDir()
	DefaultHomePath = filepath.Join(UserDir, DefaultHomeDirName)
)

func InitConfig() {
	if Home == "" {
		Home = DefaultHomePath
	}
	if err := configInit(Home); err != nil {
		panic(err)
	}
}

type (
	Config struct {
		App   App   `toml:"app"`
		Chain Chain `toml:"chain"`
		Mysql Mysql `toml:"mysql"`
		Log   Log   `toml:"log"`
	}

	Mysql struct {
		MysqlStr     string `toml:"mysql_str"`
		DBLog        uint8  `toml:"db_log"`
		DatabaseName string `toml:"database_name"`
	}

	Log struct {
		LogFileName     string `toml:"log_file_name"`
		LogmaxAge       int    `toml:"logmax_age"`
		LogrotationTime int64  `toml:"logrotation_time"`
	}

	Chain struct {
		Source ChainCfg `toml:"source"`
		Dest   ChainCfg `toml:"dest"`
	}

	ChainCfg struct {
		Cache      Cache      `toml:"cache"`
		Tendermint Tendermint `toml:"tendermint"`
		Eth        Eth        `toml:"eth"`
		Bsc        Bsc        `toml:"bsc"`
		ChainType  string     `toml:"chain_type"`
	}

	// bsc config ============================================================

	Bsc struct {
		URI                   string       `toml:"uri"`
		ChainID               uint64       `toml:"chain_id"`
		ChainName             string       `toml:"chain_name"`
		Contracts             EthContracts `toml:"contracts"`
		UpdateClientFrequency uint64       `toml:"update_client_frequency"`
		GasLimit              uint64       `toml:"gas_limit"`
		MaxGasPrice           uint64       `toml:"max_gas_price"`
		CommentSlot           int64        `toml:"comment_slot"`
		TipCoefficient        float64      `toml:"tip_coefficient"`
	}

	// eth config ============================================================

	Eth struct {
		URI                   string       `toml:"uri"`
		ChainID               uint64       `toml:"chain_id"`
		ChainName             string       `toml:"chain_name"`
		Contracts             EthContracts `toml:"contracts"`
		UpdateClientFrequency uint64       `toml:"update_client_frequency"`
		GasLimit              uint64       `toml:"gas_limit"`
		MaxGasPrice           uint64       `toml:"max_gas_price"`
		CommentSlot           int64        `toml:"comment_slot"`
		TipCoefficient        float64      `toml:"tip_coefficient"`
	}

	EthContracts struct {
		Packet    EthContractCfg `toml:"packet"`
		AckPacket EthContractCfg `toml:"ack_packet"`
		Client    EthContractCfg `toml:"client"`
	}

	EthContractCfg struct {
		Addr       string `toml:"addr"`
		Topic      string `toml:"topic"`
		OptPrivKey string `toml:"opt_priv_key"`
	}

	// Tendermit config =====================================================

	Tendermint struct {
		ChainName             string   `toml:"chain_name"`
		ChainID               string   `toml:"chain_id"`
		RPCAddr               string   `toml:"rpc_addr"`
		GrpcAddr              string   `toml:"grpc_addr"`
		Gas                   uint64   `toml:"gas"`
		Key                   ChainKey `toml:"key"`
		Fee                   Fee      `toml:"fee"`
		RequestTimeout        uint     `toml:"request_timeout"` //TODO no use
		UpdateClientFrequency uint64   `toml:"update_client_frequency"`
	}

	Fee struct {
		Denom  string `toml:"denom"`
		Amount int64  `toml:"amount"`
	}

	ChainKey struct {
		Name     string `toml:"name"`
		Password string `toml:"password"`
		Mnemonic string `toml:"mnemonic"`
	}

	// =====================================================================

	App struct {
		MetricAddr   string   `toml:"metric_addr"`
		Env          string   `toml:"env"`
		LogLevel     string   `toml:"log_level"`
		ChannelTypes []string `toml:"channel_types"`
	}

	Cache struct {
		Filename    string `toml:"filename"`
		StartHeight uint64 `toml:"start_height"`
	}
)

func LoadConfigs() *Config {
	if Home == "" {
		Home = DefaultHomePath
	}
	if LocalConfig == "" {
		LocalConfig = filepath.Join(Home, DefaultConfigDirName, DefaultConfigName)
	}
	cfg := Config{}
	tools.InitTomlConfigs([]*tools.ConfigMap{
		{
			FilePath: LocalConfig,
			Pointer:  &cfg,
		},
	})
	return &cfg
}

func configInit(home string) error {
	cfgDir := path.Join(home, DefaultConfigDirName)
	cfgPath := path.Join(cfgDir, DefaultConfigName)
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		if _, err := os.Stat(home); os.IsNotExist(err) {
			if err = os.Mkdir(home, os.ModePerm); err != nil {
				return err
			}
		}
		if err = os.Mkdir(cfgDir, os.ModePerm); err != nil {
			return err
		}
	}
	f, err := os.Create(cfgPath)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(defaultConfig()); err != nil {
		return err
	}
	return nil
}

func defaultConfig() string {
	return defaultCfg
}
