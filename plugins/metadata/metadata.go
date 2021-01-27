package metadata

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ClientValueCallWrap(kv ...string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		pairs := metadata.Pairs(kv...)
		outgoingContext := metadata.NewOutgoingContext(ctx, pairs)
		return invoker(outgoingContext, method, req, reply, cc, opts...)
	}
}
