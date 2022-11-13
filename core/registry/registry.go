package registry

import (
	"context"
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
