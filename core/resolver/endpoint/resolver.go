package endpoint

import (
	"context"
	"github.com/liujunren93/share/core/registry"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"
)

const Name = "shareEndpoint"

func init() {

	resolver.Register(&endpointBuilder{})
}

type endpointBuilder struct{}

func (e *endpointBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	var address []resolver.Address

	service, err := registry.RegistryInstance.GetService(target.Endpoint)
	if err != nil {
		return nil, err
	}
	up := func(serviceList []*registry.Service) {
		for _, s := range serviceList {
			values := new(attributes.Attributes).WithValues("weight", s.Weight)
			address = append(address, resolver.Address{Addr: s.Endpoint,Attributes: values})
		}

		cc.UpdateState(resolver.State{
			Addresses: address,

		})

	}

	go registry.RegistryInstance.Watch(target.Endpoint, context.TODO(), up)
	up(service)


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
