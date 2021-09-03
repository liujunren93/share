package client

import (
	"context"
	"github.com/liujunren93/share/core/balancer/roundRobin"
	"github.com/liujunren93/share/core/registry"
	"google.golang.org/grpc"
	"strings"
	"time"
)

type options struct {
	name string
	serverOpt
	callWrappers []grpc.UnaryClientInterceptor
	ctx          context.Context
	registry     registry.Registry
	balancer     string
	timeout      time.Duration // 请求超时
}
type serverOpt struct {
	namespace string
}

var DefaultOptions = options{

	ctx: context.TODO(),
	balancer:roundRobin.Name,
	serverOpt: serverOpt{
		namespace: "go/micro/srv",
	},
	timeout: time.Second * 3,
}

func WithTimeout(timeout time.Duration) option {
	return func(o *options) {
		o.timeout = timeout
	}
}

func WithNamespace(namespace string) option {
	return func(o *options) {
		o.namespace = strings.Replace(namespace, ".", "/", -1)
	}
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

//func WithGrpcOpts(ops ...grpc.DialOption) option {
//	ops = append(ops, grpc.WithInsecure())
//	return func(o *options) {
//		o.grpcOpts = ops
//	}
//}
func WithCallWrappers(ops ...grpc.UnaryClientInterceptor) option {
	return func(o *options) {
		o.callWrappers = append(o.callWrappers, ops...)
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
