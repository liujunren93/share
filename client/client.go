package client

import (
	"context"
	"github.com/liujunren93/share/client/selector"
	"github.com/liujunren93/share/log"
	"google.golang.org/grpc"
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

	if c.options.ctx == nil {
		c.options.ctx = context.TODO()
	}

}

func (c *Client) Client(serverName string) (*grpc.ClientConn, error) {
	service, err := c.options.registry.GetService(serverName)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	if len(service) == 0 {
		log.Logger.Error("service not found")
		return nil, nil
	}
	round := selector.Round(service)
	s := round()

	return grpc.DialContext(c.options.ctx, s, c.options.callOption...)
}
