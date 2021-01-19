package main

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/example/proto"
	"github.com/liujunren93/share/plugins/opentrace"
	"github.com/liujunren93/share/plugins/validator"
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share_utils/wrapper/openTrace"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

type hello struct {
}

func (h hello) Say(ctx context.Context, req *proto.Req) (*proto.Res, error) {

	var res proto.Res
	res.Msg = req.Name + ":hello world1"
	return &res, nil
}

func main() {
	newJaeger, _, err := openTrace.NewJaeger("app", "127.0.0.1:6831")
	fmt.Println(err)
	opentracing.SetGlobalTracer(newJaeger)

	grpcServer := server.NewGrpcServer(
		server.WithName("app"),
		//server.WithAddress("127.0.0.1:2222"),
		server.WithHdlrWrappers(validator.NewHandlerWrapper(),
			opentrace.ServerGrpcWrap(newJaeger),

		),
	)

	r ,err:= etcd.NewRegistry(registry.WithAddrs("127.0.0.1:2379"))
	if err != nil {
		panic(err)
	}
	grpcServer.Registry(r)
	proto.RegisterHelloWorldServer(grpcServer.Server().(*grpc.Server), new(hello))

	grpcServer.Run()
}

//func NewJaeger(servicename string, addr string) (opentracing.Tracer, io.Closer, error) {
//
//	cfg := config.Configuration{
//		ServiceName: servicename,
//		Sampler: &config.SamplerConfig{
//			Type:  jaeger.SamplerTypeConst,
//			Param: 1,
//		},
//
//		Reporter: &config.ReporterConfig{
//			LogSpans:            true,
//			BufferFlushInterval: 1 * time.Second,
//		},
//	}
//
//	sender, err := jaeger.NewUDPTransport(addr, 0)
//	if err != nil {
//		fmt.Println(err)
//		return nil, nil, err
//	}
//
//	reporter := jaeger.NewRemoteReporter(sender)
//	// Initialize tracer with a logger and a metrics factory
//	tracer, closer, err := cfg.NewTracer(
//		config.Reporter(reporter),
//	)
//
//	return tracer, closer, err
//}
