package client

import (
	"github.com/liujunren93/share/core/balancer/weight"
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

func (c *Client) Dial(serverName string) (*grpc.ClientConn, error) {

	opts := c.options.grpcOpts

	opts = append(opts, grpc.WithBalancerName(weight.Name))

	return grpc.Dial(BuildDirectTarget(serverName), opts...)
}
