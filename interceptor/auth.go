package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

func authServerUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//TODO: add auth
	return handler(ctx, req)
}

func authServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	//TODO: add auth
	return handler(srv, ss)
}
