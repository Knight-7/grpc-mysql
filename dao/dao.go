package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DAO struct {
	db *sqlx.DB
}

func NewDAO(db *sqlx.DB) *DAO {
	return &DAO{
		db: db,
	}
}
