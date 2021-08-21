package client

import (
	"context"
	"google.golang.org/grpc"
)

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
