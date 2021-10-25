package weight

import (
	"errors"
	thisBalancer "github.com/liujunren93/share/core/balancer"
	"github.com/liujunren93/share/utils"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"sync"
)

const Name = "share_weight"

func newBuilder() balancer.Builder {
	return base.NewBalancerBuilder(Name, &weightPickerBuilder{}, base.Config{HealthCheck: false})
}
func init() {

	balancer.Register(newBuilder())
}

type weightPickerBuilder struct {
}

func (w *weightPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {

	var scs []thisBalancer.SubConn
	for sc, val := range info.ReadySCs {

		value := val.Address.Attributes.Value("weight")
		scs = append(scs, thisBalancer.SubConn{
			SubConn: sc,
			Weight:  value.(int),
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

	var conn balancer.SubConn
	switch len(p.subConns) {
	case 0:
		p.mu.Unlock()
		return balancer.PickResult{},errors.New("service not found")
	case 1:
		conn=p.subConns[0].SubConn
	default:
		conn = choose(p.subConns).SubConn

	}
	p.mu.Unlock()
	return balancer.PickResult{SubConn: conn}, nil
}

func choose(subConns []thisBalancer.SubConn) thisBalancer.SubConn {
	totalWeight:=0
	for _, subConn := range subConns {
		totalWeight+=subConn.Weight
	}
	randInt := utils.RandInt(totalWeight)

	var begin, end, maxWeight int
	var defaultSubconn thisBalancer.SubConn
	for _, subconn := range subConns {
		if subconn.Weight > maxWeight {
			maxWeight = subconn.Weight
			defaultSubconn = subconn
		}
		end += subconn.Weight
		if randInt >= begin && randInt < end {
			return subconn
		}
		begin += subconn.Weight
	}
	return defaultSubconn
}
