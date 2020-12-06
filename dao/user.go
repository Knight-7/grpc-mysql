package dao

import (
	"context"
	"rpc-mysql/pkg/models"
	pb "rpc-mysql/pkg/proto"
	"rpc-mysql/pkg/utils"
)

func (d *DAO) GetUser(ctx context.Context, id int) (*pb.User, error) {
	var users []*models.User
	sql := `
		SELECT id, name, age, email
		FROM user
		WHERE id = ?
	`
	err := ExecSelectSQL(ctx, d.db, &users, sql, id)
	if err != nil {
		return nil, err
	}

	user := users[0]
	return &pb.User{
		Id:    int32(user.Id),
		Name:  user.Name,
		Age:   int32(user.Age),
		Email: user.Email,
	}, nil
}

func (d *DAO) GetUsers(ctx context.Context) ([]*pb.User, error) {
	var users []*models.User
	sql := `SELECT id, name, age, email FROM user`
	err := ExecSelectSQL(ctx, d.db, &users, sql)
	if err != nil {
		return nil, err
	}
	return utils.ToRpcUsers(users), nil
}

func (d DAO) AddUser(ctx context.Context, user *pb.User) (int, error) {
	sql := `
		INSERT INTO user (name, age, email)
		VALUES (?, ?, ?)
	`
	id, err := ExecInsertSQL(ctx, d.db, sql, user.Name, user.Age, user.Email)
	if err != nil {
		return -1, nil
	}

	return id, nil
}

func (d *DAO) UpdateUser(ctx context.Context, user *pb.User) (int, error) {
	sql := `
		UPDATE user SET name = ?, age = ?, email = ?
		WHERE id = ?
	`
	rows, err := ExecUpdateSQL(ctx, d.db, sql, user.Name, user.Age, user.Email, user.Id)
	if err != nil {
		return -1, err
	}

	return rows, nil
}

func (d *DAO) DeleteUser(ctx context.Context, id int) (int, error) {
	sql := `DELETE FROM user WHERE id = ?`
	rows, err := ExecDeleteSQL(ctx, d.db, sql, id)
	if err != nil {
		return -1, nil
	}

	return rows, nil
}
