package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func authServerUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Printf("method: %s, authority: %s\n", info.FullMethod, "knight")
	return handler(ctx, req)
}

func authServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("method %s, authority: %s\n", info.FullMethod, "knight")
	return handler(srv, ss)
}
