package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/example/proto"

	"math/rand"
	"strconv"
	"time"

	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share/wrapper/validator"
)

type hello struct {
}

func (h hello) Say(ctx context.Context, req *proto.Req) (*proto.Res, error) {
	var res proto.Res
	fmt.Println(1111)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res.Msg = req.Name + ":hello world" + strconv.FormatInt(r.Int63n(1000), 10)
	return &res, nil
}

var weight *int

func init() {
	weight = flag.Int("w", 10, "")
	flag.Parse()
}
func main() {

	grpcServer := server.NewGrpcServer(
		server.WithNamespace("aaaaaa"),
		server.WithName("test"),

		server.WithHdlrWrappers(validator.NewHandlerWrapper()),
	)
	r, err := etcd.NewRegistry(registry.WithAddrs("http://node1:2379", "http://node1:3379", "http://node1:4379"))
	if err != nil {
		panic(err)
	}

	grpcServer.Registry(r, registry.WithWeight(*weight))
	proto.RegisterHelloWorldServer(grpcServer.Server(), new(hello))

	grpcServer.Run()
}
