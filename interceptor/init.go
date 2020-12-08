package interceptor

import (
	"github.com/sirupsen/logrus"
	"rpc-mysql/pkg/clientset"
	"rpc-mysql/pkg/config"
)

var interceptorLogger *logrus.Logger

func SetupAuthAndLog(cfg *config.Config) {
	var err error
	tmpLog := logrus.New()
	tmpLog.SetReportCaller(true)

	interceptorLogger, err = clientset.NewLogger(cfg)
	if err != nil {
		tmpLog.Fatalln(err)
	}
}
