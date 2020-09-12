package server

import (
	"fmt"
	"github.com/liujunren93/share/log"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/utils"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type maxMsgSizeKey struct{}

type Option func(*Options)

var (
	// DefaultMaxMsgSize define maximum message size that server can send
	// or receive.  Default value is 4MB.
	DefaultMaxMsgSize = 1024 * 1024 * 4
)

type grpcServer struct {
	srv     *grpc.Server
	options *Options
}

func (g *grpcServer) getMaxMsgSize() int {
	if g.options.Ctx == nil {
		return DefaultMaxMsgSize
	}
	s, ok := g.options.Ctx.Value(maxMsgSizeKey{}).(int)
	if !ok {
		return DefaultMaxMsgSize
	}
	return s
}

func NewGrpcServer(options ...Option) *grpcServer {
	var s grpcServer
	s.options = &defaultOptions
	s.Init(options...)
	return &s

}

func (g *grpcServer) Init(options ...Option) {
	for _, o := range options {
		o(g.options)
	}
	gopts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(g.getMaxMsgSize()),
		grpc.MaxSendMsgSize(g.getMaxMsgSize()),
	}

	for _, wrapper := range g.options.HdlrWrappers {
		gopts = append(gopts, wrapper)
	}
	g.srv = grpc.NewServer(gopts...)
}

func (g *grpcServer) Registry(reg registry.Registry) error {
	if g.options.Name == "" {
		log.Logger.Panicln("service name cannot be empty")
		//return errors.New("service name cannot be empty")
	}
	ser := registry.Service{
		Name:     g.options.Name,
		Version:  g.options.Version,
		Node:     utils.GetUuidV3(reg.GetPrefix()),
		Endpoint: g.options.Address.addr,
	}

	return reg.Registry(&ser)
}

func (g *grpcServer) Server() *grpc.Server {
	return g.srv
}

func (g *grpcServer) Run() error {
	listen, err := net.Listen("tcp", g.options.Address.addr)
	if err != nil {
		log.Logger.Panic(err)
	}

	go func() {
		if err := g.srv.Serve(listen); err != nil {
			log.Logger.Errorf("[share] Server [grpc] error:%s \n", err)
			os.Exit(0)
		}

	}()
	fmt.Printf("[share] Server [grpc] Listening on %s \n", listen.Addr().String())
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, Shutdown()...)

	select {
	// wait on kill signal
	case <-ch:
	// wait on context cancel

	}
	return nil
}

func Shutdown() []os.Signal {
	return []os.Signal{
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL,
	}
}
