package main

import (
	"context"
	"github.com/liujunren93/share/example/proto"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/registry/etcd"
	"github.com/liujunren93/share/server"
	"google.golang.org/grpc"
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
	proto.RegisterHelloWorldServer(grpcServer.Server().(*grpc.Server), new(hello))

	grpcServer.Run()
}
