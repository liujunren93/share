package registry

import "context"

type Registry interface {
	Init(...Option) error
	Registry(*Service, ...Option) error
	GetService(serverName string, opts ...Option) (*[]*Service, error)
	Watch(serverName string, ctx context.Context, srvList *[]*Service)
	GetPrefix() string
}

type Service struct {
	Name     string            `json:"name"`
	Version  string            `json:"version"`
	Metadata map[string]string `json:"metadata"`
	Node     string            `json:"node"`
	Endpoint string            `json:"endpoint"`
}

type Option func(*Options)
