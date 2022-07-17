package endpoint

import (
	"context"
	"time"

	"github.com/liujunren93/share/core/registry"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"
)

const Name = "shareresolver"

func init() {

	resolver.Register(&endpointBuilder{})
}

type endpointBuilder struct{}

func (e *endpointBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
	service, err := registry.RegistryInstance.GetService(ctx, target.Endpoint)
	if err != nil {
		return nil, err
	}
	up := func(serviceList []*registry.Service) {
		var address []resolver.Address
		for _, s := range serviceList {
			values := attributes.New("weight", s.Weight)
			address = append(address, resolver.Address{Addr: s.Endpoint, Attributes: values})
		}
		cc.UpdateState(resolver.State{
			Addresses: address,
		})
	}

	registry.RegistryInstance.RegistryMonitor(target.Endpoint, up)
	if len(service) != 0 {
		up(service)
	}
	return &shareResolver{}, nil
}

func (e *endpointBuilder) Scheme() string {
	return Name
}

type shareResolver struct {

	//ccc resolver.ClientConn
}

func (s *shareResolver) ResolveNow(options resolver.ResolveNowOptions) {
}

func (s *shareResolver) Close() {

}
