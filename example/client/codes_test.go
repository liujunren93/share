package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/codes/json"
	"github.com/liujunren93/share/core/balancer/roundRobin"
	"github.com/liujunren93/share/example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCodes(t *testing.T) {

	// r, err := etcd.NewRegistry(registry.WithAddrs("http://node1:2379", "http://node1:3379", "http://node1:4379"))
	// if err != nil {
	// 	panic(err)
	// }
	newClient := client.NewClient(client.WithBuildTargetFunc(func(args ...string) string {
		return "127.0.0.1:9090"

	}), client.WithBalancer(roundRobin.Name), client.WithTimeout(time.Hour), client.WithNamespace("aaaaaa"), client.WithGrpcDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())))

	conn, err := newClient.Client("test")

	if err != nil {
		panic(err)
	}
	newClient.AddCallOptions(grpc.CallContentSubtype(json.Name))
	for {

		var res = new(proto.Res)
		fmt.Println(222)
		// var data interface{}
		err := newClient.Invoke(context.TODO(), "proto.helloWorld/say", []byte(`{"name":"aaa"}`), res, conn)
		fmt.Println(err, res)
		// // // fmt.Println(runtime.NumGoroutine())
		// mathClient := proto.NewHelloWorldClient(conn)
		// add, err := mathClient.Say(context.TODO(), &proto.Req{Name: "test"})
		// fmt.Println(add, err)
	}
}
