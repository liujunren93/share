package client

import (
	"fmt"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/balancer/roundrobin"
)

type option func(*options)

type Client struct {
	options *options
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
	opts := c.options.grpcOpts
	opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy":"%s"}`, c.options.balancer)))
	return grpc.Dial(BuildDirectTarget(serverName), opts...)
}
