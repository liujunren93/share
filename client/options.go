package client

import (
	"context"
	"github.com/liujunren93/share/client/selector"
	"github.com/liujunren93/share/registry"
	"google.golang.org/grpc"
)

type options struct {
	name         string
	serverName   string
	callWrappers []grpc.UnaryClientInterceptor
	grpcOpts     []grpc.DialOption
	ctx          context.Context
	Selector     selector.RoundType
	registry     registry.Registry
}

var DefaultOptions = options{
	grpcOpts: []grpc.DialOption{grpc.WithInsecure()},
	ctx:      context.TODO(),
	Selector: selector.RoundRobin,
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

func WithSelector(roundType selector.RoundType) option {
	return func(o *options) {
		o.Selector = roundType
	}
}

func WithRegistry(r registry.Registry) option {

	return func(o *options) {
		o.registry = r
	}
}
