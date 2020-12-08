package clientset

import (
	"github.com/sirupsen/logrus"
	"os"
	"rpc-mysql/pkg/config"
)

func NewLogger(cfg *config.Config) (*logrus.Logger, error) {
	log := logrus.New()

	level, err := logrus.ParseLevel(cfg.GetLogLevel())
	if err != nil {
		return nil, err
	}

	formatter := &logrus.JSONFormatter{
		TimestampFormat:  cfg.GetLogTimeFormatter(),
		DisableTimestamp: cfg.GetLogDisableTimestamp(),
	}

	log.SetLevel(level)
	logrus.SetFormatter(formatter)

	file, err := os.OpenFile(cfg.GetLogFilePath(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)

	return log, nil
}
