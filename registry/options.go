package registry

import (
	"context"
	"crypto/tls"
	"time"
)

type Options struct {
	Addrs     []string
	Prefix    string
	Timeout   time.Duration
	TLSConfig *tls.Config
}

type WatchOptions struct {
	Ctx     context.Context
	SrvList *[]*Service
}

var DefaultOptions = Options{
	Addrs:     []string{"127.0.0.1:2379"},
	Prefix:    "share",
	Timeout:   time.Second * 2,
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

func WithTimeout(duration time.Duration) Option {
	return func(options *Options) {
		options.Timeout = duration
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
