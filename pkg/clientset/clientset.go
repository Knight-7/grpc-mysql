package clientset

import (
	"github.com/jmoiron/sqlx"
	"rpc-mysql/pkg/config"
)

type ClientsSet struct {
	MySQL *sqlx.DB
}

func NewClientset(cfg *config.Config) (*ClientsSet, error) {
	mysql, err := NewMySQLClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ClientsSet{
		MySQL: mysql,
	}, nil
}
