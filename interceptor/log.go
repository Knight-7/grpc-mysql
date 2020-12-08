package interceptor

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
)

func logServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	logrus.Info(info.FullMethod)
	return handler(ctx, req)
}

func logServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("request method: %s\n", info.FullMethod)
	logrus.Info(info.FullMethod)
	return handler(srv, ss)
}
