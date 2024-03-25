package k8s

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/liujunren93/share/core/registry"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type k8sRegistry struct {
	options      *registry.Options
	k8sClientset *kubernetes.Clientset
	namespace    string
	serverList   sync.Map
	k8sOptions   Option
	monitors     sync.Map
}

type Option struct {
	ServicePortName string
}

func NewRegistry(option Option, options ...registry.Option) (*k8sRegistry, error) {

	k := k8sRegistry{
		options:    &registry.DefaultOptions,
		k8sOptions: option,
	}
	err := k.init(options...)
	if err != nil {
		return nil, err
	}
	registry.RegistryInstance = &k
	return &k, nil

}

func (r *k8sRegistry) init(options ...registry.Option) error {
	for _, o := range options {
		o(r.options)
	}
	conf, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(conf)
	if err != nil {
		return err
	}
	r.k8sClientset = clientset

	return nil
}

func (r *k8sRegistry) Registry(ctx context.Context, svc *registry.Service) error {
	r.namespace = svc.Namespace
	// endpoints, err := r.k8sClientset.CoreV1().Endpoints(r.namespace).Get(ctx, svc.Name, v1.GetOptions{})
	// if err != nil {
	// 	return err
	// }
	// if svc.Weight > 0 {
	// 	r.k8sClientset.CoreV1().Endpoints(r.namespace).Apply(ctx, &corev1.EndpointsApplyConfiguration{}, v1.ApplyOptions{})

	// }
	return nil
}
func (r *k8sRegistry) GetService(ctx context.Context, serverName string) ([]*registry.Service, error) {
	// registry.GetServicePath(prefix string, srvName string)
	data, ok := r.serverList.Load(serverName)
	if ok {
		list, ok := data.([]*registry.Service)
		if ok {
			return list, nil
		}

	}
	endpoints, err := r.k8sClientset.CoreV1().Endpoints(r.namespace).Get(ctx, serverName, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	for _, v := range endpoints.Subsets {
		var port int32
		for _, pt := range v.Ports {
			if pt.Name == r.k8sOptions.ServicePortName {
				port = pt.Port
			}
		}
		for _, addr := range v.Addresses {
			r.serverList.Store(serverName, &registry.Service{
				Name:      serverName,
				Node:      "",
				Endpoint:  addr.IP + ":" + strconv.Itoa(int(port)),
				Weight:    0,
				Namespace: r.namespace,
			})
		}

	}

	return nil, nil
}

func (r *k8sRegistry) Watch(ctx context.Context, serverName string) {
	wc, err := r.k8sClientset.CoreV1().Endpoints(r.namespace).Watch(ctx, v1.ListOptions{})
	if err != nil {
		return
	}
	for {
		select {
		case event := <-wc.ResultChan():
			fmt.Println(event)
		}
	}

}
func (r *k8sRegistry) GetPrefix() string {
	return r.options.Prefix
}
func (r *k8sRegistry) RegistryMonitor(serverName string, f func([]*registry.Service)) {
	if _, loaded := r.monitors.LoadOrStore(serverName, f); !loaded {
		go r.Watch(context.TODO(), serverName)
	}
}
func (*k8sRegistry) UnRegistry(*registry.Service) error
