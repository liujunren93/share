package client

import (
	"context"
	"strings"
	"time"

	"github.com/liujunren93/share/core/balancer/roundRobin"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/wrapper"
	"google.golang.org/grpc"
)

type options struct {
	name string
	serverOpt
	callWrappers    map[string]grpc.UnaryClientInterceptor
	ctx             context.Context
	registry        registry.Registry
	balancer        string
	timeout         time.Duration // 请求超时
	buildTargetFunc BuildTargetFunc
	grpcOpts        []grpc.DialOption
}
type serverOpt struct {
	namespace string
}

var DefaultOptions = options{
	ctx:      context.TODO(),
	balancer: roundRobin.Name,
	serverOpt: serverOpt{
		namespace: "share/srv",
	},
	timeout:      time.Second * 3,
	callWrappers: make(map[string]grpc.UnaryClientInterceptor),
}

func WithTimeout(timeout time.Duration) OptionFunc {
	return func(o *options) {
		o.timeout = timeout
	}
}
func WithGrpcDialOption(opts ...grpc.DialOption) OptionFunc {
	return func(o *options) {
		o.grpcOpts = append(o.grpcOpts, opts...)
	}
}

func WithBuildTargetFunc(buildTargetFunc BuildTargetFunc) OptionFunc {
	return func(o *options) {
		o.buildTargetFunc = buildTargetFunc
	}
}

func WithNamespace(namespace string) OptionFunc {
	return func(o *options) {
		o.namespace = strings.Replace(namespace, ".", "/", -1)
	}
}

func WithBalancer(name string) OptionFunc {
	return func(o *options) {
		o.balancer = name
	}
}
func WithName(name string) OptionFunc {
	return func(o *options) {
		o.name = name
	}
}

func WithCallWrappers(wraps ...wrapper.CallWrapper) OptionFunc {
	return func(o *options) {
		for _, wrap := range wraps {
			interceptor, name := wrap()
			o.callWrappers[name] = interceptor
		}
	}
}

func WithCtx(ctx context.Context) OptionFunc {
	return func(o *options) {
		o.ctx = ctx
	}
}

func WithRegistry(r registry.Registry) OptionFunc {

	return func(o *options) {
		o.registry = r
	}
}
