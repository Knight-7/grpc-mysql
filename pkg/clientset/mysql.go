package clientset

import (
	"github.com/jmoiron/sqlx"
	"rpc-mysql/pkg/config"
	"time"
)

func NewMySQLClient(cfg *config.Config) (*sqlx.DB, error) {
	mysqlDsn := cfg.GetMySQLDsn()

	db, err := sqlx.Connect("mysql", mysqlDsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MySQL.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MySQL.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.MySQL.MaxLifetime) * time.Second)

	return db, nil
}
