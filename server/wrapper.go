package server

import (
	"context"
	"google.golang.org/grpc"
)

func UnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.ServerOption {
	n := len(interceptors)
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInter(currentCtx, currentReq, info, currentHandler)
			}
		}
		chainedHandler := handler
		for i := 0; i < n; i++ {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}

		return chainedHandler(ctx, req)

	})
}

func UnaryClient(interceptors ...grpc.UnaryClientInterceptor) grpc.DialOption {
	n := len(interceptors)
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		chainer := func(currentInter grpc.UnaryClientInterceptor, currentInvoker grpc.UnaryInvoker) grpc.UnaryInvoker {
			return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
				return currentInter(ctx, method, req, reply, cc, currentInvoker, opts...)
			}
		}
		chainedInvoker := invoker
		for i := 0; i < n; i++ {
			chainedInvoker = chainer(interceptors[i], chainedInvoker)
		}
		return chainedInvoker(ctx, method, req, reply, cc, opts...)
	})
}
