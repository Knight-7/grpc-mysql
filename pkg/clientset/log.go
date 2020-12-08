package clientset

import (
	"github.com/sirupsen/logrus"
	"os"
	"rpc-mysql/pkg/config"
)

func InitLog(cfg *config.Config) error {
	level, err := logrus.ParseLevel(cfg.GetLogLevel())
	if err != nil {
		return err
	}

	formatter := &logrus.JSONFormatter{
		TimestampFormat:  cfg.GetLogTimeFormatter(),
		DisableTimestamp: cfg.GetLogDisableTimestamp(),
	}

	logrus.SetLevel(level)

	logrus.SetFormatter(formatter)

	file, err := os.OpenFile(cfg.GetLogFilePath(), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	logrus.SetOutput(file)

	return nil
}
