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

	if !isDirExist(cfg.GetLogFilePath()) {
		err = os.MkdirAll(cfg.GetLogFilePath(), os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(cfg.GetLogFilePath()+"info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)

	return log, nil
}

func isDirExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
	}
	return false
}
