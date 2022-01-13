package logger

import (
	log "github.com/sirupsen/logrus"

	"github.com/teleport-network/teleport-relayer/app/config"
)

func Logger(cfg *config.Config) *log.Logger {
	logger := log.New()
	if cfg.App.Env == "prod" {
		logger.SetFormatter(&log.JSONFormatter{})
	} else {
		logger.SetFormatter(&log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})
	}
	switch cfg.App.LogLevel {
	case "debug":
		logger.SetLevel(log.DebugLevel)
	case "error":
		logger.SetLevel(log.ErrorLevel)
	case "warn":
		logger.SetLevel(log.WarnLevel)
	default:
		logger.SetLevel(log.InfoLevel)
	}
	return logger
}
