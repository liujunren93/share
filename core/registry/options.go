package registry

import (
	"context"
	"crypto/tls"
)

type Options struct {
	Addrs        []string
	Prefix       string
	RegistryCtx  context.Context
	GetServerCtx context.Context
	TLSConfig    *tls.Config
}

type WatchOptions struct {
	Ctx     context.Context
	SrvList *[]*Service
}

var DefaultOptions = Options{
	Addrs:     []string{"127.0.0.1:2379"},
	Prefix:    "registry",
	TLSConfig: nil,
}

func WithAddrs(addrs ...string) Option {
	return func(options *Options) {
		options.Addrs = addrs
	}
}

func WithPrefix(prefix string) Option {
	return func(options *Options) {
		options.Prefix = prefix
	}
}

func WithRegistryCtx(ctx context.Context) Option {
	return func(options *Options) {
		options.RegistryCtx = ctx
	}
}
func WithGetServerCtx(ctx context.Context) Option {
	return func(options *Options) {
		options.GetServerCtx = ctx
	}
}

func WithTLSConfig(tls *tls.Config) Option {
	return func(options *Options) {
		options.TLSConfig = tls
	}
}

//func WithWatchCtx(ctx context.Context) WatchOption {
//	return func(options *WatchOptions) {
//		options.Ctx = ctx
//	}
//}
//
//func WithWatchSrvList(list *[]*Service) WatchOption {
//	return func(options *WatchOptions) {
//		options.SrvList = list
//	}
//}
