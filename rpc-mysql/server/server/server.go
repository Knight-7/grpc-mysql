package server

import "google.golang.org/grpc"

type RPCServer struct {
	Server *grpc.Server
	Addr   string
}

func NewRPCServer(addr string) *RPCServer {
	rpcServer := grpc.NewServer()
	return &RPCServer{
		Server: rpcServer,
		Addr:   addr,
	}
}
