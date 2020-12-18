package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io"
	"log"
	"rpc-mysql/pkg/auth"
	"rpc-mysql/pkg/config"
	pb "rpc-mysql/pkg/proto"
)

func getUser(ctx context.Context, client pb.DAOClient) {
	fmt.Println("start get users")
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
	fmt.Println()
}

func getUserById(ctx context.Context, client pb.DAOClient, id int) {
	fmt.Println("start get user by id")
	resp, err := client.GetUserById(ctx, &pb.GetUserByIdRequest{
		Id: int32(id),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*resp.User)
	fmt.Println()
}

func addUser(ctx context.Context, client pb.DAOClient, user *pb.User) {
	fmt.Println("start add user")
	req := &pb.AddUserRequest{User: user}
	resp, err := client.AddUser(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Message)
	fmt.Println()
}

func updateUser(ctx context.Context, client pb.DAOClient, user *pb.User) {
	fmt.Println("start update user")
	req := &pb.UpdateUserRequest{
		User: user,
	}
	resp, err := client.UpdateUser(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Message)
	fmt.Println()
}

func deleteUser(ctx context.Context, client pb.DAOClient, id int) {
	fmt.Println("start delete user")
	req := &pb.DeleteUserRequest{
		Id: int32(id),
	}

	resp, err := client.DeleteUser(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Message)
}

func main() {
	filePath := flag.String("config", "config.yaml", "get config path")
	flag.Parse()

	err := config.LoadYAMLConfig(*filePath)
	if err != nil {
		log.Fatalln(err)
	}

	cfg := config.GetConfig()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithPerRPCCredentials(&auth.Authorizer{
		Login:    "login",
		Password: "pass",
		OpenTLS:  true,
	}))

	creds, err := auth.GetClientCreds(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	opts = append(opts, creds)

	conn, err := grpc.Dial(cfg.GetClientTarget(), opts...)
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewDAOClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	getUser(ctx, client)

	/*addUser(ctx, client, &pb.User{
		Name:  "lin",
		Age:   23,
		Email: "lin@lin.com",
	})

	updateUser(ctx, client, &pb.User{
		Id:    3,
		Name:  "haoyouking",
		Age:   22,
		Email: "haoyouking@haoyouking.com",
	})

	deleteUser(ctx, client, 6)*/

	updateUser(ctx, client, &pb.User {
		Id:    4,
		Name:  "yujian.ou",
		Age:   23,
		Email: "yujian.ou@123.com",
	})

	getUserById(ctx, client, 1)

	getUser(ctx, client)
}
