package registry

import (
	"context"
	"crypto/tls"
)

type Options struct {
	Addrs     []string
	Prefix    string
	ctx       context.Context
	Lease     int64
	TLSConfig *tls.Config
}

type WatchOptions struct {
	Ctx     context.Context
	SrvList *[]*Service
}

var DefaultOptions = Options{
	Addrs:     []string{"127.0.0.1:2379"},
	Prefix:    "registry",
	TLSConfig: nil,
	Lease:     5,
}

func WithAddrs(addrs ...string) Option {
	return func(options *Options) {
		options.Addrs = addrs
	}
}
func WithLease(lease int64) Option {
	return func(options *Options) {
		options.Lease = lease
	}
}

func WithPrefix(prefix string) Option {
	return func(options *Options) {
		options.Prefix = prefix
	}
}

func WithCtx(ctx context.Context) Option {
	return func(options *Options) {
		options.ctx = ctx
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
