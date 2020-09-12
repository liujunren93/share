package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func CallOption() grpc.DialOption {
	return grpc.WithUnaryInterceptor(ClientInterceptor())
}

func ClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string,req, reply interface{}, cc *grpc.ClientConn,invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Println(req)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
