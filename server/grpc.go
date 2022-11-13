package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/log"
	"github.com/liujunren93/share/utils"
	recover2 "github.com/liujunren93/share/wrapper/recover"
	"google.golang.org/grpc"
)

type maxMsgSizeKey struct{}

var (
	// DefaultMaxMsgSize define maximum message size that server can send
	// or receive.  Default value is 4MB.
	DefaultMaxMsgSize = 1024 * 1024 * 4
)

type GrpcServer struct {
	srv      *grpc.Server
	options  *Options
	listener net.Listener
	stopList []func()
}

func (g *GrpcServer) getMaxMsgSize() int {
	if g.options.Ctx == nil {
		return DefaultMaxMsgSize
	}
	s, ok := g.options.Ctx.Value(maxMsgSizeKey{}).(int)
	if !ok {
		return DefaultMaxMsgSize
	}
	return s
}

func NewGrpcServer(options ...Option) *GrpcServer {
	var s GrpcServer
	opt := defaultOptions
	s.options = &opt
	s.init(options)
	return &s

}

func (g *GrpcServer) init(options []Option) {
	for _, o := range options {
		o(g.options)
	}
	if g.options.Mode != "debug" {
		g.options.HandleWrappers = append(g.options.HandleWrappers, recover2.NewServerWrapper())
	}
	if g.listener == nil {
		listen, err := net.Listen("tcp", g.options.ListenAddr)
		if err != nil {
			log.Logger.Panic(err)
		}
		//m := cmux.New(listen)
		//grpcL := m.MatchWithWriters(
		//	cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"),
		//)

		g.options.ListenAddr = listen.Addr().String()
		g.listener = listen
	}
	gopts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(g.getMaxMsgSize()),
		grpc.MaxSendMsgSize(g.getMaxMsgSize()),
	}

	it := UnaryServer(g.options.HandleWrappers...)
	gopts = append(gopts, it)
	gopts = append(gopts, g.options.GrpcOpts...)

	g.srv = grpc.NewServer(gopts...)
}

func (g *GrpcServer) Registry(reg registry.Registry, servers ...registry.Server) error {

	if g.options.Name == "" {
		log.Logger.Panicln("service name cannot be empty")
		return errors.New("service name cannot be empty")
	}
	ip, err := utils.GetIntranetIp()
	if err != nil {
		return err
	}
	endpoint := strings.Replace(g.options.ListenAddr, "[::]", ip.String(), 1)

	ser := registry.Service{
		Name:      g.options.Name,
		Version:   g.options.Version,
		Node:      utils.GetUuidV3(reg.GetPrefix()),
		Endpoint:  endpoint,
		Namespace: g.options.Namespace,
	}
	for _, server := range servers {
		server(&ser)
	}
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
	err = reg.Registry(ctx, &ser)
	if err != nil {
		return err
	}
	g.stopList = append(g.stopList, func() {
		reg.UnRegistry(&ser)
	})
	return nil
}

func (g *GrpcServer) Server() grpc.ServiceRegistrar {
	return g.srv
}

func (g *GrpcServer) Run() error {
	go func() {
		if err := g.srv.Serve(g.listener); err != nil {
			log.Logger.Errorf("[share] Server [grpc] error:%s \n", err)
			os.Exit(0)
		}
	}()
	fmt.Printf("[share] Server [grpc] Listening on %s \n", g.listener.Addr().String())
	g.WatchSignal()
	return nil
}

func (g *GrpcServer) WatchSignal() {
	ch := make(chan os.Signal, 1)
	signals := []os.Signal{
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL,
	}
	signal.Notify(ch, signals...)
	select {
	// wait on kill signal
	case <-ch:
		g.srv.Stop()
		for _, stop := range g.stopList {
			stop()
		}
		fmt.Printf("[share] Server [grpc] stop:%s", g.options.Name)
		// wait on context cancel

	}

}

func (g *GrpcServer) RegistryStopFunc(f func()) {
	g.stopList = append(g.stopList, f)
}
