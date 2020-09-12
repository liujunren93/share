# share
基于grpc的一个微服务框架
<ul>
<li>
服务发现
</li>
<li>
服务注册
</li>
<li>
grpc 服务封装
</li>
<li>
grpc 客户端封装 
</li>
</ul>
server:<br>

```golang

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

```

client:

```golang

func main() {
	r := etcd.NewRegistry()
	r.Init(registry.WithAddrs("127.0.0.1:2379"))
	newClient := client.NewClient()
	newClient.Init(client.WithRegistry(r),client.WithCallOption(CallOption()))
	conn, err := newClient.Client("app")

	mathClient := proto.NewHelloWorldClient(conn)
	add, err := mathClient.Say(context.TODO(), &proto.Req{
		Name: "test",
	})
	fmt.Println(add, err)
}

```# share
