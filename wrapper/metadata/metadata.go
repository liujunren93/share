package metadata

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func NewClientWrapper(kv ...string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}

		pairs(md,kv...)
		outgoingContext := metadata.NewOutgoingContext(ctx, md)
		return invoker(outgoingContext, method, req, reply, cc, opts...)
	}
}

func pairs(md metadata.MD,kv ...string) metadata.MD {
	if len(kv)%2 == 1 {
		panic(fmt.Sprintf("metadata: Pairs got the odd number of input pairs for metadata: %d", len(kv)))
	}

	var key string
	for i, s := range kv {
		if i%2 == 0 {
			key = strings.ToLower(s)
			continue
		}
		md[key] = append(md[key], s)
	}
	return md
}