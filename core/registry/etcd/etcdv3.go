package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/log"
	"go.etcd.io/etcd/clientv3"
	"strings"
	"sync"
	"time"
)

type etcdRegistry struct {
	client          *clientv3.Client
	options         *registry.Options
	serverList      sync.Map
	monitors        sync.Map
	serverNamespace string
}

func NewRegistry(options ...registry.Option) (*etcdRegistry, error) {

	e := etcdRegistry{
		options: &registry.DefaultOptions,
	}
	err := e.init(options...)
	if err != nil {
		return nil, err
	}
	registry.RegistryInstance = &e
	return &e, nil

}

func (e *etcdRegistry) init(options ...registry.Option) error {
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

func (e *etcdRegistry) Registry(service *registry.Service) error {

	lease := clientv3.NewLease(e.client)

	ctx := e.options.RegistryCtx
	if ctx == nil {
		ctx = context.TODO()
	}
	grant, err := lease.Grant(ctx, 5)
	if err != nil {
		return err
	}
	lease.KeepAlive(ctx, grant.ID)
	e.serverNamespace=e.options.Prefix+service.Namespace
	_, err = e.client.Put(ctx, RegisterPath(e.options.Prefix, service), encode(service), clientv3.WithLease(grant.ID))

	fmt.Printf("[share] Registering on [etcd]:%s  \n", RegisterPath(e.options.Prefix, service))
	fmt.Printf("[share] Registering name: %s  \n", service.Name)
	return err
}

func (e *etcdRegistry) GetService(serverName string) ([]*registry.Service, error) {

	load, ok := e.serverList.Load(serverName)
	if ok {
		return load.([]*registry.Service), nil
	}
	ctx := e.options.GetServerCtx
	if ctx == nil {
		ctx, _ = context.WithTimeout(context.TODO(), time.Second*2)
	}

	get, err := e.client.Get(ctx, GetServicePath(e.options.Prefix, serverName), clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}

	var serviceList []*registry.Service
	for _, kv := range get.Kvs {

		r := decode(kv.Value)
		if r.Namespace+"/"+r.Name == serverName {
			serviceList = append(serviceList, r)
		}

	}

	e.serverList.Store(serverName, serviceList)

	return serviceList, nil
}

//RegistryMonitor 注册监视器
func (e *etcdRegistry) RegistryMonitor(serverName string, f func([]*registry.Service)) {

	if _, loaded := e.monitors.LoadOrStore(serverName, f); !loaded {
		go e.Watch(serverName, context.TODO())
	}

}

func (e *etcdRegistry) Watch(serverName string, ctx context.Context) {

	watch := e.client.Watch(ctx, GetServicePath(e.options.Prefix, serverName), clientv3.WithPrefix(), clientv3.WithPrevKV())
	for response := range watch {
		var serviceList []*registry.Service
		load, ok := e.serverList.Load(serverName)
		if ok && load != nil {
			serviceList = load.([]*registry.Service)
		}
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

					serviceList = append(serviceList, service)
				}
				if event.IsModify() {
					for i, service := range serviceList {
						if "node_"+service.Node == node[len(node)-1] {
							serviceList[i] = service
						}
					}
				}

			case clientv3.EventTypeDelete:
				for i := 0; i < len(serviceList); i++ {
					if "node_"+serviceList[i].Node == node[len(node)-1] {
						serviceList = append(serviceList[:i], serviceList[i+1:]...)
						i--
					}
				}
			}
		}
		e.serverList.Store(serverName, serviceList)
		if value, ok := e.monitors.Load(serverName); ok {
			if f, ok := value.(func([]*registry.Service)); ok {
				f(serviceList)
			}
		}
	}
}

func inSlice(list []string, s string) bool {
	for _, str := range list {
		if str == s {
			return true
		}
	}
	return false
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

func RegisterPath(prefix string, srv *registry.Service) string {
	return fmt.Sprintf("/%s/%s/%s/node_%s", prefix, srv.Namespace, srv.Name, srv.Node)
}

func GetServicePath(prefix, srvName string) string {
	return fmt.Sprintf("/%s/%s", prefix, srvName)
}
