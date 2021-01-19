package server

import (
	"github.com/liujunren93/share/core/registry"
)

type Server interface {
	Init(options ...Option)
	Registry(reg registry.Registry) error
	Server() interface{}
	Run() error
}

type Option func(*Options)
