package timeout

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

type timeout struct {
}

func NewClientWrapper(duration time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, _ = context.WithTimeout(ctx, duration)
		return invoker(ctx,method,req,reply,cc,opts...)
	}
}
