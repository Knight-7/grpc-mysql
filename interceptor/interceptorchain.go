package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

func unaryServerInterceptChain(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	//获取拦截器的长度
	l := len(interceptors)
	//如下我们返回一个拦截器
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//在这个拦截器中，我们做一些操作
		//构造一个链
		chain := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInter(
					currentCtx,
					currentReq,
					info,
					currentHandler)
			}
		}
		//声明一个handler
		chainHandler := handler
		for i := l - 1; i >= 0; i-- {
			//递归一层一层调用
			chainHandler = chain(interceptors[i], chainHandler)
		}
		//返回结果
		return chainHandler(ctx, req)
	}
}

func streamServerInterceptorChain(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	l := len(interceptors)
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		chain := func(currentInter grpc.StreamServerInterceptor, currentHandle grpc.StreamHandler) grpc.StreamHandler {
			return func(currentSrv interface{}, currentSS grpc.ServerStream) error {
				return currentInter(
					currentSrv,
					currentSS,
					info,
					currentHandle)
			}
		}
		chainHandler := handler
		for i := l - 1; i >= 0; i-- {
			chainHandler = chain(interceptors[i], chainHandler)
		}
		return chainHandler(srv, ss)
	}
}

func NewServerUnaryInterceptor() grpc.UnaryServerInterceptor {
	var interceptors []grpc.UnaryServerInterceptor
	interceptors = append(interceptors, logServerInterceptor)
	return unaryServerInterceptChain(interceptors...)
}

func NewServerStreamInterceptor() grpc.StreamServerInterceptor {
	var interceptors []grpc.StreamServerInterceptor
	interceptors = append(interceptors, logServerStreamInterceptor)
	return streamServerInterceptorChain(interceptors...)
}
