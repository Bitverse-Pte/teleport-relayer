package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/teleport-network/teleport-relayer/version"
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
	logger.Infof("appName:%s",version.Name)
	logger.Infof("appCommit:%s",version.Commit)
	channelMap := channels.NewChannelMap(cfg, logger)
	return &App{
		channelMap: channelMap,
		logger:     logger,
	}
}

func (a *App) Start() {
	s := gocron.NewScheduler(time.UTC)
	r := gin.New()
	r.Use(gin.Recovery())
	for chainName, channel := range a.channelMap {
		a.logger.Printf("relay packet for %s\n", chainName)
		channel.RelayTask(s)
		r.PUT(fmt.Sprintf("/relayer/%v/height", chainName), channel.UpgradeRelayHeight)
		r.GET(fmt.Sprintf("/relayer/%v/height", chainName), channel.ViewRelayHeight)
		r.PUT(fmt.Sprintf("/relayer/%v/extra_wait", chainName), channel.UpgradeExtraWait)
		r.GET(fmt.Sprintf("/relayer/%v/extra_wait", chainName), channel.ViewExtraWait)
	}
	s.StartAsync()
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func (a *App) EvmClientSync() {
	s := gocron.NewScheduler(time.UTC)
	for chainName, channel := range a.channelMap {
		a.logger.Printf("relay packet for %s\n", chainName)
		channel.EvmClientUpdate(s)
	}
	s.StartBlocking()
}
