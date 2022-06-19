package server

import (
	"context"
	"strings"

	"google.golang.org/grpc"
)

type Options struct {
	Name           string
	Address        string
	Namespace      string
	Version        string
	Mode           string //debubg
	Ctx            context.Context
	GrpcOpts       []grpc.ServerOption
	HandleWrappers []grpc.UnaryServerInterceptor
}

type Option func(*Options)

var defaultOptions = Options{
	Address:        ":0",
	Namespace:      "go/share/srv",
	Version:        "latest",
	Mode:           "release",
	Ctx:            context.TODO(),
	HandleWrappers: []grpc.UnaryServerInterceptor{},
}

//mode release,debug
func WithMode(mode string) Option {
	return func(o *Options) {
		o.Mode = mode
	}
}

func WithName(name string) Option {
	return func(options *Options) {
		options.Name = strings.Replace(name, ".", "/", -1)
	}
}

func WithAddress(addr string) Option {
	return func(options *Options) {
		options.Address = addr
	}
}

func WithNamespace(namespace string) Option {
	return func(options *Options) {
		options.Namespace = strings.Replace(namespace, ".", "/", -1)
	}
}

func WithCtx(ctx context.Context) Option {
	return func(options *Options) {
		options.Ctx = ctx
	}
}

func WithHdlrWrappers(wrapper ...grpc.UnaryServerInterceptor) Option {
	return func(options *Options) {
		options.HandleWrappers = wrapper
	}
}

func WithGrpcOpts(option ...grpc.ServerOption) Option {
	return func(options *Options) {
		options.GrpcOpts = option
	}
}
