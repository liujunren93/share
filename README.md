# share
基于grpc的一个微服务框架
<ul>
<li>
服务发现
</li>
<li>
服务注册
</li>
<li>
grpc 服务封装
</li>
<li>
grpc 客户端封装 
</li>
</ul>
server:<br>

```golang
package main

import (
	"context"
	"flag"
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
var weight *int
func init() {
	weight = flag.Int("w", 10, "")
	flag.Parse()
}
func main() {
	newJaeger, _, _ := openTrace.NewJaeger("app", "127.0.0.1:6831")
	opentracing.SetGlobalTracer(newJaeger)
	grpcServer := server.NewGrpcServer(
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
	fmt.Println(weight)
	grpcServer.Registry(r, registry.WithWeight(*weight))
	proto.RegisterHelloWorldServer(grpcServer.Server().(*grpc.Server), new(hello))
	grpcServer.Run()
}


```

client:

```golang
package main

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/example/proto"
	"github.com/liujunren93/share_utils/wrapper/openTrace"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/balancer/roundrobin"
)

func main() {

	newJaeger, _, _ := openTrace.NewJaeger("client", "127.0.0.1:6831")
	opentracing.SetGlobalTracer(newJaeger)
	r,_ := etcd.NewRegistry(registry.WithAddrs("127.0.0.1:2379"))
	newClient := client.NewClient(client.WithRegistry(r),client.WithBalancer(roundrobin.Name))

	conn, _ := newClient.Client("app")
	for {
		fmt.Scanln()
		mathClient := proto.NewHelloWorldClient(conn)

		add, err := mathClient.Say(context.TODO(), &proto.Req{
			Name: "adsa",
		})
		fmt.Println(add, err)
	}
}



``` 
