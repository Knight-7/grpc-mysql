package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

func logServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	interceptorLogger.WithFields(
		map[string]interface{}{
			"method": info.FullMethod,
		},
	).Info()
	return handler(ctx, req)
}

func logServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	interceptorLogger.WithFields(
		map[string]interface{}{
			"method": info.FullMethod,
		},
	).Info()
	return handler(srv, ss)
}
