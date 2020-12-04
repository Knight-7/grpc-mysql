package dao

import (
	"context"
	"database/sql"
	"fmt"
)

type CURD interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func ExecSelectSQL(ctx context.Context, db CURD, dest interface{}, query string, args ...interface{}) error {
	err := db.SelectContext(ctx, dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func ExecInsertSQL(ctx context.Context, db CURD, query string, args ...interface{}) (int, error) {
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(lastId), nil
}

func ExecUpdateSQL(ctx context.Context, db CURD, query string, args ...interface{}) (int, error) {
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	if rows == 0 {
		return -1, fmt.Errorf("no rows updated")
	}

	return int(rows), nil
}

func ExecDeleteSQL(ctx context.Context, db CURD, query string, args ...interface{}) (int, error) {
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	if rows == 0 {
		return -1, fmt.Errorf("no record deleted")
	}
	return int(rows), nil
}
