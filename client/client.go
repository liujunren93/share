package client

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/client/selector"
	"github.com/liujunren93/share/log"
	"github.com/liujunren93/share/serrors"
	"github.com/liujunren93/share/server"
)

type option func(*options)

type Client struct {
	options *options
}

func NewClient() *Client {
	var c Client
	c.options = &DefaultOptions
	return &c
}

func (c *Client) Init(opts ...option) {
	for _, o := range opts {
		o(c.options)
	}

}

func (c *Client) Client(serverName string) (selector.Next, error) {
	service, err := c.options.registry.GetService(serverName)
	if err != nil {
		log.Logger.Error(err)
		return nil, serrors.NotFound(err)
	}

	if *service == nil {
		log.Logger.Errorf("[share]service:%s not found", serverName)
		return nil, serrors.NotFound(fmt.Errorf("[share]service:%s not found", serverName))
	}

	c.options.grpcOpts = append(c.options.grpcOpts, server.UnaryClient(c.options.callWrappers...))
	ctx := context.WithValue(c.options.ctx, "serverName", serverName)
	return c.options.Selector(service, ctx, c.options.grpcOpts...), nil

}
