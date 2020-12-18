package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func logServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		interceptorLogger.Errorf("log meta data not found")
		return nil, status.Errorf(codes.NotFound, "log meta data not found")
	}

	var user string
	if val, ok := md["login"]; ok {
		user = val[0]
	}

	interceptorLogger.WithFields(
		map[string]interface{}{
			"method": info.FullMethod,
			"user":   user,
		},
	).Info()
	return handler(ctx, req)
}

func logServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		interceptorLogger.Errorf("log meta data not found")
		return status.Errorf(codes.NotFound, "log meta data not found")
	}

	var user string
	if val, ok := md["login"]; ok {
		user = val[0]
	}

	interceptorLogger.WithFields(
		map[string]interface{}{
			"method": info.FullMethod,
			"user": user,
		},
	).Info()
	return handler(srv, ss)
}
