package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

func interceptChain(intercepts ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	//获取拦截器的长度
	l := len(intercepts)
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
			chainHandler = chain(intercepts[i], chainHandler)
		}
		//返回结果
		return chainHandler(ctx, req)
	}
}

func NewIntercept() grpc.UnaryServerInterceptor {
	var intercepts []grpc.UnaryServerInterceptor
	intercepts = append(intercepts, logInterceptor)
	return interceptChain(intercepts...)
}
