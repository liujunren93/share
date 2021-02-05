package client

import (
	"fmt"
	"google.golang.org/grpc"
	"sync"

	//"google.golang.org/grpc/balancer/roundrobin"
)

type option func(*options)

type Client struct {
	options   *options
	endpoints sync.Map
}

func NewClient(opts ...option) *Client {
	var c Client
	c.options = &DefaultOptions
	for _, o := range opts {
		o(c.options)
	}
	return &c

}

//AddOptions
func (c *Client) AddOptions(opts ...option) {
	for _, o := range opts {
		o(c.options)
	}
}

//Client
func (c *Client) Client(serverName string) (*grpc.ClientConn, error) {
	if load, ok := c.endpoints.Load(serverName); ok {
		return load.(*grpc.ClientConn), nil
	} else {
		opts := c.options.grpcOpts
		opts = append(opts, UnaryClient(c.options.callWrappers...))
		opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy":"%s"}`, c.options.balancer)))
		if dial, err := grpc.Dial(BuildDirectTarget(c.options.namespace, serverName), opts...); err != nil {
			return nil, err
		} else {
			c.endpoints.Store(serverName, dial)
			return dial, nil
		}
	}
}
