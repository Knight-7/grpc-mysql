package rpc

import (
	"context"
	pb "rpc-mysql/pkg/proto"
	"time"
)

func (d DaoRPC) GetUsers(req *pb.GetUsersRequest, stream pb.DAO_GetUsersServer) error {
	users, err := d.d.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if err = stream.Send(&pb.GetUserResponse{User: user}); err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func (d DaoRPC) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserResponse, error) {
	user, err := d.d.GetUser(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		User: user,
	}, nil
}

func (d DaoRPC) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.ExecSQLResponse, error) {
	id, err := d.d.AddUser(ctx, req.User)
	if err != nil {
		return nil, err
	}

	return &pb.ExecSQLResponse{
		Message:      "Add user success",
		LastInsertId: int32(id),
	}, nil
}

func (d DaoRPC) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.ExecSQLResponse, error) {
	rows, err := d.d.UpdateUser(ctx, req.User)
	if err != nil {
		return nil, err
	}

	resp := &pb.ExecSQLResponse{AffectRows: int32(rows)}
	if rows == 0 {
		resp.Message = "no record updated"
	} else if rows > 1 {
		resp.Message = "tow many rows updated"
	} else if rows == 1 {
		resp.Message = "update success"
	}

	return resp, nil
}

func (d DaoRPC) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.ExecSQLResponse, error) {
	rows, err := d.d.DeleteUser(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	resp := &pb.ExecSQLResponse{AffectRows: int32(rows)}
	if rows == 0 {
		resp.Message = "no record deleted"
	} else if rows > 1 {
		resp.Message = "to many rows deleted"
	} else if rows == 1 {
		resp.Message = "delete success"
	}

	return resp, nil
}
