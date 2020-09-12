package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var defaultConfig = clientv3.Config{
	Endpoints:            []string{"http://127.0.0.1:2179"},
	AutoSyncInterval:     30 * time.Second,
	DialTimeout:          3 * time.Second,
	DialKeepAliveTime:    10 * time.Second,
	DialKeepAliveTimeout: 3 * time.Second,
	Context:              context.TODO(),
}

func DefaultConfig() clientv3.Config {
	return defaultConfig
}

func NewEtcdClient(config *clientv3.Config) (*clientv3.Client, error) {
	if config == nil {
		config = &defaultConfig
	}
	return clientv3.New(*config)
}
