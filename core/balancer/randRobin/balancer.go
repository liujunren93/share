package randRobin

import (
	thisBalancer "github.com/liujunren93/share/core/balancer"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"sync"
)

const Name = "round_robin"

func newBuilder() balancer.Builder {
	return base.NewBalancerBuilderV2(Name, &weightPickerBuilder{}, base.Config{HealthCheck: false})
}
func init() {

	balancer.Register(newBuilder())
}

type weightPickerBuilder struct {
}

func (w *weightPickerBuilder) Build(info base.PickerBuildInfo) balancer.V2Picker {

	var scs []thisBalancer.SubConn
	for sc, _ := range info.ReadySCs {
		//value := val.Address.Attributes.Value("weight")
		scs = append(scs, thisBalancer.SubConn{
			SubConn: sc,
			//Weight:  value.(int8),
		})
	}
	return &weightPickerPicker{subConns: scs}
}

type weightPickerPicker struct {
	subConns []thisBalancer.SubConn
	mu       sync.Mutex
	next     int
}

func (p *weightPickerPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {

	p.mu.Lock()
	sc := p.subConns[p.next]
	p.next = (p.next + 1) % len(p.subConns)
	p.mu.Unlock()
	return balancer.PickResult{SubConn: sc.SubConn}, nil
}
