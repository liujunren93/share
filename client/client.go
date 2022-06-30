package client

import (
	"errors"
	"fmt"
	"sync"

	"google.golang.org/grpc"
)

type OptionFunc func(*options)

type Client struct {
	options   *options
	endpoints sync.Map
}
type BuildTargetFunc func(args ...string) string

func NewClient(opts ...OptionFunc) *Client {
	var c Client
	c.options = &DefaultOptions
	for _, o := range opts {
		o(c.options)
	}
	return &c

}

//AddOptions
func (c *Client) AddOptions(opts ...OptionFunc) {
	for _, o := range opts {
		o(c.options)
	}
}

func (c *Client) buildGrpcOptions() []grpc.DialOption {
	opts := c.options.grpcOpts
	opts = append(opts, grpc.WithTimeout(c.options.timeout))

	if c.options.balancer != "" {
		opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy":"%s"}`, c.options.balancer)))
	}
	for _, v := range c.options.callWrappers {
		opts = append(opts, UnaryClient(v))
	}

	return opts
}

//Client
func (c *Client) Client(serverName string) (grpc.ClientConnInterface, error) {
	if load, ok := c.endpoints.Load(serverName); ok {
		return load.(grpc.ClientConnInterface), nil
	} else {
		opts := c.buildGrpcOptions()
		var target string

		if c.options.buildTargetFunc == nil {
			if serverName == "" {
				return nil, errors.New("serverName can not be empty")
			}
			target = defaultDirectTarget(c.options.namespace, serverName)
		} else {
			target = c.options.buildTargetFunc(c.options.namespace)
		}

		if dial, err := grpc.Dial(target, opts...); err != nil {
			return nil, err
		} else {
			c.endpoints.Store(serverName, dial)
			return dial, nil
		}
	}
}
