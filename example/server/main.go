package main

import (
	"context"
	"flag"
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
	//app.RegisterAppServer(grpcServer.Server(), new(hello))
	app.RegisterAppServer(grpcServer.Server(), new(appstr))
	grpcServer.Run()
}

type appstr string

func (a appstr) Create(ctx context.Context, req *app.CreateReq) (*app.DefaultRes, error) {
	var res app.DefaultRes
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res.Msg =   ":hello world"+strconv.FormatInt(r.Int63n(1000),10)
	return &res, nil
}

func (a appstr) Update(ctx context.Context, req *app.UpdateReq) (*app.DefaultRes, error) {
	panic("implement me")
}

func (a appstr) List(ctx context.Context, req *app.AppListReq) (*app.AppListRes, error) {
	panic("implement me")
}

func (a appstr) Info(ctx context.Context, req *app.AppInfoReq) (*app.AppInfoRes, error) {
	panic("implement me")
}

func (a appstr) Delete(ctx context.Context, req *app.DeleteReq) (*app.DefaultRes, error) {
	panic("implement me")
}
