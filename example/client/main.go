package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/core/balancer/roundRobin"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/example/proto"
)

func main() {

	r, _ := etcd.NewRegistry(registry.WithAddrs("http://127.0.0.1:2379", "http://127.0.0.1:3379", "http://127.0.0.1:4379"))
	newClient := client.NewClient(client.WithRegistry(r), client.WithBalancer(roundRobin.Name), client.WithNamespace("aaaaaa"))
	conn, err := newClient.Client("test")
	if err != nil {
		panic(err)
	}

	for {

		fmt.Scanln()
		fmt.Println(runtime.NumGoroutine())
		mathClient := proto.NewHelloWorldClient(conn)
		add, err := mathClient.Say(context.TODO(), &proto.Req{Name: "test"})
		fmt.Println(add, err)
	}
}
