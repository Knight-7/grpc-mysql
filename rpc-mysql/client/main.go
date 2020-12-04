package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io"
	pb "rpc-mysql/proto"
	"time"
)

const (
	addr = "42.192.11.222:3306"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewDAOClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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
