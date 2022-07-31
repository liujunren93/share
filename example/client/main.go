package main

import (
	"context"
	"fmt"
	"time"

	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/core/balancer/roundRobin"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/example/proto"
	"github.com/liujunren93/share/wrapper/timeout"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	r, err := etcd.NewRegistry(registry.WithAddrs("http://node1:2379", "http://node1:3379", "http://node1:4379"))
	if err != nil {
		panic(err)
	}
	newClient := client.NewClient(client.WithRegistry(r), client.WithBalancer(roundRobin.Name), client.WithNamespace("aaaaaa"), client.WithGrpcDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())), client.WithCallWrappers(timeout.NewClientWrapper(time.Second)))
	conn, err := newClient.Client("test")
	if err != nil {
		panic(err)
	}

	for {

		fmt.Scanln()
		var res = new(proto.Res)
		fmt.Println(222)
		// var data interface{}
		err := newClient.Invoke(context.TODO(), "proto.helloWorld/say", &proto.Req{Name: "test"}, res, conn)
		fmt.Println(err, res)
		// // // fmt.Println(runtime.NumGoroutine())
		// mathClient := proto.NewHelloWorldClient(conn)
		// add, err := mathClient.Say(context.TODO(), &proto.Req{Name: "test"})
		// fmt.Println(add, err)
	}
}
