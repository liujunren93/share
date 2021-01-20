package server

import (
	"context"
	"google.golang.org/grpc"
)

type address struct {
	addr      string
	IsDefault bool
}
type Options struct {
	Name           string
	Address        string
	Namespace      string
	Version        string
	Ctx            context.Context
	GrpcOpts       []grpc.ServerOption
	HandleWrappers []grpc.UnaryServerInterceptor
}

var defaultOptions = Options{
	Address:   ":0",
	Namespace: "share",
	Version:   "latest",
	Ctx:       context.TODO(),
}

func WithName(name string) Option {
	return func(options *Options) {
		options.Name = name
		//options.Name = strings.Replace(name, ".", "/", -1)
	}
}

func WithAddress(addr string) Option {
	return func(options *Options) {
		options.Address = addr
	}
}

func WithNamespace(namespace string) Option {
	return func(options *Options) {
		options.Namespace = namespace
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
