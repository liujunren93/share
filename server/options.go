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
	Name         string
	Address      address
	Namespace    string
	Version      string
	Ctx          context.Context
	HdlrWrappers []grpc.ServerOption
}

var defaultOptions = Options{
	Address: address{
		addr:      ":0",
		IsDefault: true,
	},
	Namespace: "share",
	Version:   "latest",
	Ctx:       context.TODO(),
}

func WithName(name string) Option {
	return func(options *Options) {
		options.Name = name
	}
}

func WithAddress(addr string) Option {
	return func(options *Options) {
		options.Address = address{
			addr:      addr,
			IsDefault: false,
		}
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

func WithHdlrWrappers(wrapper ...grpc.ServerOption) Option {
	return func(options *Options) {
		options.HdlrWrappers = wrapper
	}
}
