package registry

import (
	"context"
)

var RegistryInstance Registry

type Registry interface {
	Registry(*Service) error
	GetService(serverName string) ([]*Service, error)
	Watch(serverName string, ctx context.Context)
	GetPrefix() string
	RegistryMonitor(serverName string, f func([]*Service))
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
type Server func(*Service)

func WithWeight(weight int) Server {
	return func(service *Service) {
		service.Weight = weight
	}
}
