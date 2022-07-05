package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	endpoints := []string{
		"http://192.168.48.134:2379", "http://192.168.48.134:3379", "http://192.168.48.134:4379",
	}
	client3, err := clientv3.New(
		clientv3.Config{
			Endpoints: endpoints,
		},
	)
	if err != nil {
		panic(err)
	}

	// go func() {
	// 	for {
	// 		client3.Put(context.Background(), "/tttt/ttt", "11")
	// 		time.Sleep(time.Second)
	// 	}
	// }()

	go func() {
		w := client3.Watch(context.Background(), "/share/registry/aaaaaa/test/node_7f044460-474c-3e4e-b871-dd620a22e63f", clientv3.WithPrefix(), clientv3.WithPrevKV())
		for aa := range w {
			fmt.Println(aa)
		}
	}()

	time.Sleep(100 * time.Hour)

}
