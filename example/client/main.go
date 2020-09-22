package main

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/client/selector"
	"github.com/liujunren93/share/example"
	"github.com/liujunren93/share/example/proto"
	"github.com/liujunren93/share/plugins/metadata"
	"github.com/liujunren93/share/plugins/opentrace"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/registry/etcd"
	"github.com/opentracing/opentracing-go"
	metadata2 "google.golang.org/grpc/metadata"
	"time"
)

func main() {
	newJaeger, _, _ := example.NewJaeger("client", "127.0.0.1:6831")
	opentracing.SetGlobalTracer(newJaeger)
	r := etcd.NewRegistry()
	r.Init(registry.WithAddrs("127.0.0.1:2379"))
	newClient := client.NewClient()
	newClient.Init(client.WithRegistry(r),
		client.WithSelector(selector.Round),
		client.WithCallWrappers(opentrace.ClientGrpcCallWrap(newJaeger),
			metadata.ClientValueCallWrap("aaa", "bbb")),
	)
	conn, err := newClient.Client("app")
	fmt.Println(err)
	for {
		time.Sleep(time.Second)
		clientConn, err := conn()
		if err != nil {
			fmt.Println(err)
			continue
			//return
		}
		mathClient := proto.NewHelloWorldClient(clientConn)

		md := metadata2.Pairs("key", "val")
		//
		//// 新建一个有 metadata 的 context
		ctx := metadata2.NewOutgoingContext(context.Background(), md)
		add, err := mathClient.Say(ctx, &proto.Req{
			Name: "adsa",
		})
		fmt.Println(add, err)
	}

}
