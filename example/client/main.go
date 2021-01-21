package main

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/example/proto"
	"github.com/liujunren93/share_utils/wrapper/openTrace"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/balancer/roundrobin"
)

func main() {
	newJaeger, _, _ := openTrace.NewJaeger("client", "127.0.0.1:6831")
	opentracing.SetGlobalTracer(newJaeger)
	r,_ := etcd.NewRegistry(registry.WithAddrs("127.0.0.1:2379"))
	newClient := client.NewClient(client.WithRegistry(r),client.WithBalancer(roundrobin.Name),client.WithNamespace("go.micro.srv"))

	conn, err := newClient.Client("account")
	if err != nil {
		//panic(err)
	}
	for {
		fmt.Scanln()
		mathClient := proto.NewHelloWorldClient(conn)
		add, err := mathClient.Say(context.TODO(), &proto.Req{
			Name: "adsa",
		})
		fmt.Println(add, err)
	}
}



