package registry

import (
	"context"
	"fmt"
)

var RegistryInstance Registry

type Registry interface {
	Registry(context.Context, *Service) error
	GetService(ctx context.Context, serverName string) ([]*Service, error)
	Watch(ctx context.Context, serverName string)
	GetPrefix() string
	RegistryMonitor(serverName string, f func([]*Service))
	UnRegistry(*Service) error
}

type Service struct {
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Metadata  map[string]string `json:"metadata"`
	Node      string            `json:"node"`
	Endpoint  string            `json:"endpoint"`
	Weight    int               `json:"weight"`
	Namespace string            `json:"namespace"`
}

type Option func(*Options)
type ServerOpt func(*Service)

func WithWeight(weight int) ServerOpt {
	return func(service *Service) {
		service.Weight = weight
	}
}

func RegisterPath(prefix string, srv *Service) string {
	return fmt.Sprintf("/%s/%s/%s/node_%s", prefix, srv.Namespace, srv.Name, srv.Node)
}

func GetServicePath(prefix, srvName string) string {
	return fmt.Sprintf("/%s/%s", prefix, srvName)
}
