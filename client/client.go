package client

import (
	"context"
	"errors"
	"github.com/liujunren93/share/client/selector"
	"github.com/liujunren93/share/log"
	"github.com/liujunren93/share/serrors"
	"github.com/liujunren93/share/server"
	"google.golang.org/grpc"
	"time"
)

type option func(*options)

type Client struct {
	options *options
}

func NewClient() *Client {
	var c Client
	c.options = &DefaultOptions
	timeout, _ := context.WithTimeout(c.options.ctx, time.Second*2)
	c.options.ctx = timeout
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

func (c *Client) Client(serverName string) (*grpc.ClientConn, *serrors.Error) {
	service, err := c.options.registry.GetService(serverName)
	if err != nil {
		log.Logger.Error(err)
		return nil, serrors.NotFound(err, nil)
	}
	if len(service) == 0 {
		log.Logger.Error("service not found")
		return nil, serrors.NotFound(errors.New("service not found"), nil)
	}
	round := selector.Round(service)
	s := round()
	c.options.grpcOpts = append(c.options.grpcOpts, server.UnaryClient(c.options.callWrappers...))
	dialContext, err := grpc.DialContext(c.options.ctx, s, c.options.grpcOpts...)
	return dialContext, serrors.InternalServerError(err, nil)
}
