package weight

import (
	"fmt"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"sync"
)

const Name = "weighted_round_robin"

func newBuilder() balancer.Builder {
	return base.NewBalancerBuilderV2(Name, &weightPickerBuilder{}, base.Config{HealthCheck: false})
}
func init() {

	balancer.Register(newBuilder())
}

type weightPickerBuilder struct {
}

func (w *weightPickerBuilder) Build(info base.PickerBuildInfo) balancer.V2Picker {

	var scs []balancer.SubConn
	for sc := range info.ReadySCs {

		scs = append(scs, sc)
	}
	fmt.Println(1111)
	return &weightPickerPicker{subConns: scs}
}

type weightPickerPicker struct {
	subConns []balancer.SubConn
	mu       sync.Mutex
	next     int
}

func (p *weightPickerPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
fmt.Println(p.subConns)
	p.mu.Lock()
	sc := p.subConns[p.next]
	p.next = (p.next + 1) % len(p.subConns)
	p.mu.Unlock()
	return balancer.PickResult{SubConn: sc}, nil
}
