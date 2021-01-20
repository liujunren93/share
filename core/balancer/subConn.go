package balancer

import "google.golang.org/grpc/balancer"

type SubConn struct {
	SubConn balancer.SubConn
	Weight  int
}
