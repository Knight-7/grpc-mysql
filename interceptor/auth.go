package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

func authServerUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if err := authorizer.Auth(ctx); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func authServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := authorizer.Auth(ss.Context()); err != nil {
		return err
	}

	return handler(srv, ss)
}
