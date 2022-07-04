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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	r, err := etcd.NewRegistry(registry.WithAddrs("http://192.168.48.134:2379", "http://192.168.48.134:3379", "http://192.168.48.134:4379"))
	if err != nil {
		fmt.Println(err)
	}
	newClient := client.NewClient(client.WithRegistry(r), client.WithBalancer(roundRobin.Name), client.WithNamespace("aaaaaa"), client.WithGrpcDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())))
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
