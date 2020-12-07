package server

import "google.golang.org/grpc"

type RPCServer struct {
	Server *grpc.Server
	Addr   string
}

func NewRPCServer(addr string, opt ...grpc.ServerOption) *RPCServer {
	rpcServer := grpc.NewServer(opt...)
	return &RPCServer{
		Server: rpcServer,
		Addr:   addr,
	}
}
