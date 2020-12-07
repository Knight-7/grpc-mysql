package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func logInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Printf("request method: %s", info.FullMethod)
	return handler(ctx, req)
}
