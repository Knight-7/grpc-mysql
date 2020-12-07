package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func logServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Printf("request method: %s\n", info.FullMethod)
	return handler(ctx, req)
}

func logServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("request method: %s\n", info.FullMethod)
	return handler(srv, ss)
}
