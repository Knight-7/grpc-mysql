package server

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"rpc-mysql/dao"
	pb "rpc-mysql/proto"
	"time"
)

const (
	dbAddr     = "127.0.0.1:33060"
	dbUser     = "root"
	dbPasswd   = "@Fight7!"
	dbName     = "grpc_test"
	serverAddr = "0.0.0.0:3306"
)

type daoServer struct {
	d *dao.DAO
}

func (d daoServer) GetUsers(req *pb.GetUsersRequest, stream pb.DAO_GetUsersServer) error {
	users, err := d.d.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if err = stream.Send(&pb.GetUserResponse{User: user}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}

	return nil
}

func (d daoServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserResponse, error) {
	user, err := d.d.GetUser(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		User: user,
	}, nil
}

func (d daoServer) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.ExecSQLResponse, error) {
	id, err := d.d.AddUser(ctx, req.User)
	if err != nil {
		return nil, err
	}

	return &pb.ExecSQLResponse{
		Message:      "Add user success",
		LastInsertId: int32(id),
	}, nil
}

func (d daoServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.ExecSQLResponse, error) {
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

func (d daoServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.ExecSQLResponse, error) {
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

func newServer(d *dao.DAO) *grpc.Server {
	daoserver := daoServer{d: d}

	server := grpc.NewServer()
	pb.RegisterDAOServer(server, daoserver)
	return server
}

func StartServer() {
	listen, err := net.Listen("tcp", serverAddr)
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer listen.Close()

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", dbUser, dbPasswd, dbAddr, dbName))
	if err != nil {
		grpclog.Fatalln(err)
	}

	d := dao.NewDAO(db)
	rpcServer := newServer(d)

	fmt.Println("rpc server start...")
	if err = rpcServer.Serve(listen); err != nil {
		grpclog.Fatalln(err)
	}
}
