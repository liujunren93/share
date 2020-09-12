package selector

import (
	"github.com/shareChina/share/registry"
	"github.com/shareChina/share/utils"
)

type Next func() string

func Round(s []*registry.Service) Next {

	return func() string {
		randInt := utils.RandInt(len(s))
		return s[randInt].Endpoint
	}
}

func RoundRobin(s []*registry.Service) Next {
	randInt := utils.RandInt(len(s))
	return func() string {
		return s[randInt%randInt].Endpoint
	}
}
