package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)
func ServerOption() grpc.ServerOption {
	return grpc.UnaryInterceptor(aaa())
}

func aaa() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println(req)
		return handler(ctx, req)
	}
}
