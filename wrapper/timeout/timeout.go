package timeout

import (
	"context"
	"fmt"
	"time"

	"github.com/liujunren93/share/wrapper"
	"google.golang.org/grpc"
)

const CLIENT_NAME = "timeout"

func NewClientWrapper(duration time.Duration) wrapper.CallWrapper {
	return func() (grpc.UnaryClientInterceptor, string) {
		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			fmt.Println(111)
			ctx, _ = context.WithTimeout(ctx, duration)
			return invoker(ctx, method, req, reply, cc, opts...)
		}, CLIENT_NAME
	}

}
