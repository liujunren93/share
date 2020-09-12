package main

import (
	"context"
	"fmt"
	"github.com/shareChina/share/client"
	"github.com/shareChina/share/example/proto"
	"github.com/shareChina/share/registry"
	"github.com/shareChina/share/registry/etcd"
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
