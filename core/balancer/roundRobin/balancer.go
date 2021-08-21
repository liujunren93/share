package roundRobin

import (
	thisBalancer "github.com/liujunren93/share/core/balancer"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"sync"
)

const Name = "share_round_robin"

func newBuilder() balancer.Builder {
	return base.NewBalancerBuilderV2(Name, &roundRobinPickerBuilder{}, base.Config{HealthCheck: false})
}
func init() {

	balancer.Register(newBuilder())
}


type roundRobinPickerBuilder struct {
}

func (*roundRobinPickerBuilder) Build(info base.PickerBuildInfo) balancer.V2Picker {
	var scs []thisBalancer.SubConn
	for sc, _ := range info.ReadySCs {
		//value := val.Address.Attributes.Value("weight")
		scs = append(scs, thisBalancer.SubConn{
			SubConn: sc,
		})
	}
	return &roundRobinPickerPicker{subConns: scs}
}

type roundRobinPickerPicker struct {
	subConns []thisBalancer.SubConn
	mu       sync.Mutex
	next     int
}

func (p *roundRobinPickerPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	sc := p.subConns[p.next]
	p.next = (p.next + 1) % len(p.subConns)
	p.mu.Unlock()
	return balancer.PickResult{SubConn: sc.SubConn}, nil
}
