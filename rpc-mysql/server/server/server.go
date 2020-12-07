package server

import "google.golang.org/grpc"

type RPCServer struct {
	Server *grpc.Server
	Addr   string
}

func NewRPCServer(addr string, opts ...grpc.ServerOption) *RPCServer {
	rpcServer := grpc.NewServer(opts...)
	return &RPCServer{
		Server: rpcServer,
		Addr:   addr,
	}
}
