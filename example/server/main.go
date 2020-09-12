package main

import (
	"context"
	"github.com/shareChina/share/example/proto"
	"github.com/shareChina/share/registry"
	"github.com/shareChina/share/registry/etcd"
	"github.com/shareChina/share/server"
)

type hello struct {
}

func (h hello) Say(ctx context.Context, req *proto.Req) (*proto.Res, error) {

	var res proto.Res
	res.Msg = req.Name + ":hello world"

	return &res, nil
}

func main() {

	grpcServer := server.NewGrpcServer(
		server.WithName("app"),
		server.WithAddress(":2222"),
		server.WithHdlrWrappers(ServerOption()),
		)

	r := etcd.NewRegistry()
	r.Init(registry.WithAddrs("127.0.0.1:2379"))
	grpcServer.Registry(r)
	proto.RegisterHelloWorldServer(grpcServer.Server(), new(hello))

	grpcServer.Run()
}
