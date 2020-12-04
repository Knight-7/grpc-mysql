package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io"
	pb "rpc-mysql/proto"
)

const (
	addr = "42.192.11.222:3306"
)

func getUser(ctx context.Context, client pb.DAOClient) {
	req := &pb.GetUsersRequest{}
	stream, err := client.GetUsers(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(*user)
	}
}

func addUser(ctx context.Context, client pb.DAOClient, user *pb.User) {
	req := &pb.AddUserRequest{User: user}
	resp, err := client.AddUser(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Message)
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewDAOClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	getUser(ctx, client)

	addUser(ctx, client, &pb.User{
		Name:  "yujian",
		Age:   23,
		Email: "yujian@yujian.com",
	})

	getUser(ctx, client)
}
