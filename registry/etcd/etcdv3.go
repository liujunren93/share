package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/liujunren93/share/log"
	"github.com/liujunren93/share/registry"
	"strings"
	"time"
)

type etcdRegistry struct {
	client  *clientv3.Client
	options *registry.Options
}

func NewRegistry() *etcdRegistry {
	return &etcdRegistry{
		options: &registry.DefaultOptions,
	}
}

func (e *etcdRegistry) Init(options ...registry.Option) error {
	for _, o := range options {
		o(e.options)
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints:            e.options.Addrs,
		AutoSyncInterval:     30 * time.Second,
		DialTimeout:          3 * time.Second,
		DialKeepAliveTime:    10 * time.Second,
		DialKeepAliveTimeout: 3 * time.Second,
	})
	if err != nil {
		return err
	}
	e.client = client
	return nil
}

func (e *etcdRegistry) Registry(service *registry.Service, options ...registry.Option) error {
	err := e.Init(options...)
	if err != nil {
		return err
	}
	lease := clientv3.NewLease(e.client)

	ctx := e.options.RegistryCtx
	if ctx == nil {
		ctx, _ = context.WithTimeout(ctx, time.Second*2)
	}
	grant, err := lease.Grant(ctx, 5)
	if err != nil {
		return err
	}
	lease.KeepAlive(ctx, grant.ID)
	_, err = e.client.Put(context.TODO(), RegisterPath(e.options.Prefix, service), encode(service), clientv3.WithLease(grant.ID))

	fmt.Printf("[share] Registering on [etcd]:%s  \n", RegisterPath(e.options.Prefix, service))
	fmt.Printf("[share] Registering name: %s  \n", service.Name)
	return err
}

func (e *etcdRegistry) GetService(serverName string, option ...registry.Option) ([]*registry.Service, error) {
	ctx := e.options.GetServerCtx
	if ctx == nil {
		ctx, _ = context.WithTimeout(ctx, time.Second*2)
	}
	get, err := e.client.Get(ctx, GetServicePath(e.options.Prefix, serverName), clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}
	serviceMap := make(map[string]*registry.Service)
	for _, kv := range get.Kvs {
		r := decode(kv.Value)
		serviceMap[string(kv.Key)] = r
	}
	var serviceList []*registry.Service
	for _, r := range serviceMap {
		serviceList = append(serviceList, r)
	}
	return serviceList, nil
}

func (e *etcdRegistry) Watch(serverName string, ctx context.Context, srvList *[]*registry.Service) {

	watch := e.client.Watch(ctx, GetServicePath(e.options.Prefix, serverName), clientv3.WithPrefix(), clientv3.WithPrevKV())

	for response := range watch {
		if response.Err() != nil {
			log.Logger.Error(response.Err())
			return
		}
		if response.Canceled {
			log.Logger.Error(errors.New("could not get next"))
			return
		}
		for _, event := range response.Events {
			node := strings.Split(string(event.Kv.Key), "/")
			switch event.Type {
			case clientv3.EventTypePut:
				service := decode(event.Kv.Value)
				if event.IsCreate() {

					*srvList = append(*srvList, service)
				}
				if event.IsModify() {
					for i, service := range *srvList {
						if "node_"+service.Node == node[len(node)-1] {
							(*srvList)[i] = service
						}
					}
				}

			case clientv3.EventTypeDelete:
				for i, service := range *srvList {

					if "node_"+service.Node == node[len(node)-1] {
						*srvList = append((*srvList)[:i], (*srvList)[i+1:]...)
					}
				}

			}
		}
	}
}

func (e *etcdRegistry) GetPrefix() string {
	return e.options.Prefix
}

//encode 编码器
func encode(s *registry.Service) string {
	b, _ := json.Marshal(s)
	return string(b)
}

//decode 解码
func decode(ds []byte) *registry.Service {
	var s *registry.Service
	json.Unmarshal(ds, &s)
	return s
}

func RegisterPath(Prefix string, r *registry.Service) string {
	return fmt.Sprintf("/%s/%s/node_%s", Prefix, r.Name, r.Node)
}

func GetServicePath(Prefix, ServerName string) string {
	return fmt.Sprintf("/%s/%s", Prefix, ServerName)
}
