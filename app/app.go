package app

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/repostitory/database"
	"github.com/teleport-network/teleport-relayer/app/services/channels"
	"github.com/teleport-network/teleport-relayer/tools"
)

type App struct {
	channelMap map[string]channels.IChannel
	logger     *logrus.Logger
}

func NewApp() *App {
	cfg := config.LoadConfigs()
	logger := tools.NewLogrus(
		filepath.Join(config.DefaultHomePath, "log"),
		cfg.Log.LogFileName,
		time.Duration(24*cfg.Log.LogmaxAge)*time.Hour,
		time.Duration(cfg.Log.LogrotationTime)*time.Hour,
	)
	database.NewMysqlDB(cfg.Mysql)
	logger.Info("1. service init relayers ")
	channelMap := channels.NewChannelMap(cfg, logger)
	return &App{
		channelMap: channelMap,
		logger:     logger,
	}
}

func (a *App) Start() {
	for chainName, channel := range a.channelMap {
		a.logger.Printf("relay packet for %s\n", chainName)
		go channel.PacketSync()
		go channel.EthClientSync()
		go channel.BscClientSync()
		go channel.Relay()
		go channel.DeleteRelayedPacket()
	}
	signalHandler()
}

func (a *App) EthClientSync() {
	for chainName, channel := range a.channelMap {
		a.logger.Printf("relay packet for %s\n", chainName)
		go channel.EthClientSync()
	}
	signalHandler()
}

func (a *App) BscClientSync() {
	for chainName, channel := range a.channelMap {
		a.logger.Printf("relay packet for %s\n", chainName)
		go channel.BscClientSync()
	}
	signalHandler()
}

func signalHandler() {
	var ch = make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			time.Sleep(time.Second * 2)
			logrus.Infof("get a signal %s, stop the push-admin process", si.String())
			//s.Close()
			//s.Wait()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
