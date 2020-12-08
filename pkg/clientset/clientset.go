package clientset

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"rpc-mysql/pkg/config"
)

type ClientsSet struct {
	mySQL  *sqlx.DB
	logger *logrus.Logger
}

func NewClientset(cfg *config.Config) (*ClientsSet, error) {
	mysql, err := NewMySQLClient(cfg)
	if err != nil {
		return nil, err
	}

	log, err := NewLogger(cfg)
	if err != nil {
		return nil, err
	}

	return &ClientsSet{
		mySQL:  mysql,
		logger: log,
	}, nil
}

func (c *ClientsSet) GetMySQL() *sqlx.DB {
	return c.mySQL
}

func (c *ClientsSet) GetLogger() *logrus.Logger {
	return c.logger
}
