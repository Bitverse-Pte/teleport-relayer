package app

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/jobs/channels"
	"github.com/teleport-network/teleport-relayer/app/types"
	"github.com/teleport-network/teleport-relayer/app/utils"
	"github.com/teleport-network/teleport-relayer/tools"
	"github.com/teleport-network/teleport-relayer/version"
)

type App struct {
	channelMap map[string]channels.IChannel
	logger     *logrus.Logger
}

func NewApp() *App {
	cfg := config.LoadConfigs()

	// check bridge status when init
	if cfg.App.BridgeStatusApi != "" && cfg.App.BridgeEnable {
		status, err := utils.GetBridgeStatus(cfg.App.BridgeStatusApi)
		if err != nil || status != 1 {
			panic(fmt.Sprintf("connect bridge error %v ", err))
		}
	}

	logger := tools.NewLogrus(
		filepath.Join(config.Home, "log"),
		cfg.Log.LogFileName,
		time.Duration(24*cfg.Log.LogmaxAge)*time.Hour,
		time.Duration(cfg.Log.LogrotationTime)*time.Hour,
	)
	logger.Infof("appName:%s", version.Name)
	logger.Infof("appCommit:%s", version.Commit)
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
		r.POST(fmt.Sprintf("/relayer/%v/relay", chainName), channel.ManualRelayByHash)
		r.PUT(fmt.Sprintf("/relayer/%v/height", chainName), channel.UpgradeRelayHeight)
		r.GET(fmt.Sprintf("/relayer/%v/height", chainName), channel.ViewRelayHeight)
		r.GET(fmt.Sprintf("/relayer/%v/extra_wait", chainName), channel.ViewExtraWait)
	}
	r.POST("/relayer/start", func(context *gin.Context) {
		s.StartAsync()
		if s.IsRunning() {
			context.JSON(http.StatusOK, gin.H{
				"status": types.StatusActive,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status": types.StatusStopped,
			})
		}
	})
	r.POST("/relayer/stop", func(context *gin.Context) {
		s.Stop()
		if s.IsRunning() {
			context.JSON(http.StatusOK, gin.H{
				"status": types.StatusActive,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status": types.StatusStopped,
			})
		}
	})
	r.GET("/relayer/status", func(context *gin.Context) {
		if s.IsRunning() {
			context.JSON(http.StatusOK, gin.H{
				"status": types.StatusActive,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status": types.StatusStopped,
			})
		}
	})

	// start gocron jobs
	s.StartAsync()

	// start gin api
	port := config.Port
	if port == "" {
		port = "8080"
	}
	if err := r.Run(fmt.Sprintf(":%v", port)); err != nil {
		panic(fmt.Errorf("route run error:%+v", err))
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

func (a *App) ManualRelay(detail *types.PacketDetail, hash string) error {
	var channel channels.IChannel
	if a.channelMap[detail.ChainName] != nil {
		// as src chain
		channel = a.channelMap[detail.ChainName]
	} else {
		return errors.New(fmt.Sprintf("can not find chain %s in channel", detail.ChainName))
	}
	err := channel.ManualRelay(detail, hash)
	if err != nil {
		return err
	}
	return nil
}
