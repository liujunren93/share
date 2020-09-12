package main

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/example/proto"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/registry/etcd"
)

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
