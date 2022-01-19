package app

import (
	"path/filepath"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/jobs/channels"
	"github.com/teleport-network/teleport-relayer/tools"
)

type App struct {
	channelMap map[string]channels.IChannel
	logger     *logrus.Logger
}

func NewApp() *App {
	cfg := config.LoadConfigs()

	logger := tools.NewLogrus(
		filepath.Join(config.Home, "log"),
		cfg.Log.LogFileName,
		time.Duration(24*cfg.Log.LogmaxAge)*time.Hour,
		time.Duration(cfg.Log.LogrotationTime)*time.Hour,
	)
	//database.NewMysqlDB(cfg.Mysql)
	logger.Info("1. service init relayers ")
	channelMap := channels.NewChannelMap(cfg, logger)
	return &App{
		channelMap: channelMap,
		logger:     logger,
	}
}

func (a *App) Start() {
	s := gocron.NewScheduler(time.UTC)
	for chainName, channel := range a.channelMap {
		a.logger.Printf("relay packet for %s\n", chainName)
		channel.RelayTask(s)
	}
	s.StartBlocking()
}

func (a *App) EvmClientSync() {
	s := gocron.NewScheduler(time.UTC)
	for chainName, channel := range a.channelMap {
		a.logger.Printf("relay packet for %s\n", chainName)
		s.Every(10).Seconds().Do(func() {
			if err := channel.EvmClientUpdate(); err != nil {
				time.Sleep(10 * time.Second)
			}
		})
	}
	s.StartBlocking()
}
