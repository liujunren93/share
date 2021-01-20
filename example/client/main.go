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
)

func main() {

	newJaeger, _, _ := openTrace.NewJaeger("client", "127.0.0.1:6831")
	opentracing.SetGlobalTracer(newJaeger)
	r,err := etcd.NewRegistry(registry.WithAddrs("127.0.0.1:2379"))
	newClient := client.NewClient(client.WithRegistry(r))

	conn, err := newClient.Dial("app")
	fmt.Println(err)
	//os.Exit(1)
	for {
		fmt.Scanln()
		mathClient := proto.NewHelloWorldClient(conn)

		add, err := mathClient.Say(context.TODO(), &proto.Req{
			Name: "adsa",
		})
		fmt.Println(add, err)
	}



}

