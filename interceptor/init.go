package interceptor

import (
	"github.com/sirupsen/logrus"
	"rpc-mysql/pkg/auth"
	"rpc-mysql/pkg/clientset"
	"rpc-mysql/pkg/config"
)

var interceptorLogger *logrus.Logger
var authorizer *auth.Authorizer

func InitInterceptor(cfg *config.Config) {
	var err error
	tmpLog := logrus.New()
	tmpLog.SetReportCaller(true)

	interceptorLogger, err = clientset.NewLogger(cfg)
	if err != nil {
		tmpLog.Fatalln(err)
	}

	authorizer = &auth.Authorizer{
		Login:    "login",
		Password: "pass",
		OpenTLS:  false,
	}
}
