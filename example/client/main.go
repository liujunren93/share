package main

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/core/balancer/roundRobin"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/example/proto/app"
	"github.com/liujunren93/share/plugins/metadata"
	"github.com/liujunren93/share_utils/wrapper/openTrace"
	"github.com/opentracing/opentracing-go"
	"runtime"
)

func main() {

	newJaeger, _, _ := openTrace.NewJaeger("client", "127.0.0.1:6831")
	opentracing.SetGlobalTracer(newJaeger)
	r, _ := etcd.NewRegistry(registry.WithAddrs("127.0.0.1:2379"))
	newClient := client.NewClient(client.WithRegistry(r), client.WithBalancer(roundRobin.Name), client.WithNamespace("go.micro.srv"),
		client.WithCallWrappers(metadata.ClientValueCallWrap("aa", "BB")),
	)

	for {
		conn, err := newClient.Client("app")
		if err != nil {
			panic(err)
		}
		fmt.Scanln()
		fmt.Println(runtime.NumGoroutine())
		mathClient := app.NewAppClient(conn)
		add, err := mathClient.Create(context.TODO(), &app.CreateReq{
			Data:                 nil,
		})
		fmt.Println(add, err)
	}
}
