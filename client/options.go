package client

import (
	"context"
	"github.com/liujunren93/share/core/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

type options struct {
	name         string
	callWrappers []grpc.UnaryClientInterceptor
	grpcOpts     []grpc.DialOption
	ctx          context.Context
	registry     registry.Registry
	balancer     string
}

var DefaultOptions = options{
	grpcOpts: []grpc.DialOption{grpc.WithInsecure()},
	ctx:      context.TODO(),
	balancer: roundrobin.Name,
}

func WithBalancer(name string) option {
	return func(o *options) {
		o.balancer = name
	}
}
func WithName(name string) option {
	return func(o *options) {
		o.name = name
	}
}

func WithGrpcOpts(ops ...grpc.DialOption) option {
	ops = append(ops, grpc.WithInsecure())
	return func(o *options) {
		o.grpcOpts = ops
	}
}
func WithCallWrappers(ops ...grpc.UnaryClientInterceptor) option {
	return func(o *options) {
		o.callWrappers = ops
	}
}

func WithCtx(ctx context.Context) option {
	return func(o *options) {
		o.ctx = ctx
	}
}

func WithRegistry(r registry.Registry) option {

	return func(o *options) {
		o.registry = r
	}
}
