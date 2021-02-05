package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/example/proto"
	"github.com/liujunren93/share/example/proto/app"
	"github.com/liujunren93/share/plugins/opentrace"
	"github.com/liujunren93/share/plugins/validator"
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share_utils/wrapper/openTrace"
	"github.com/opentracing/opentracing-go"
	"math/rand"
	"strconv"
	"time"
)

type hello struct {
}

func (h hello) Say(ctx context.Context, req *proto.Req) (*proto.Res, error) {
	var res proto.Res

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res.Msg = req.Name + ":hello world"+strconv.FormatInt(r.Int63n(1000),10)
	return &res, nil
}

var weight *int

func init() {
	weight = flag.Int("w", 10, "")
	flag.Parse()
}
func main() {
	if true {
		go func() {
			fmt.Println(111)
		}()
	}

	newJaeger, _, _ := openTrace.NewJaeger("app", "127.0.0.1:6831")
	opentracing.SetGlobalTracer(newJaeger)
	grpcServer := server.NewGrpcServer(
		server.WithNamespace("go.micro.srv"),
		server.WithName("app"),
		//server.WithAddress("127.0.0.1:2222"),
		server.WithHdlrWrappers(validator.NewHandlerWrapper(),
			opentrace.ServerGrpcWrap(newJaeger),
		),
	)
	r, err := etcd.NewRegistry(registry.WithAddrs("127.0.0.1:2379"))
	if err != nil {
		panic(err)
	}

	grpcServer.Registry(r, registry.WithWeight(*weight))
	proto.RegisterHelloWorldServer(grpcServer.Server(), new(hello))

	grpcServer.Run()
}

