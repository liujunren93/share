package server

import (
	"errors"
	"fmt"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/log"
	"github.com/liujunren93/share/utils"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
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
	s.options = &defaultOptions
	s.init(options)
	return &s

}

func (g *GrpcServer) init(options []Option) {
	for _, o := range options {
		o(g.options)
	}
	if g.options.Mode == "debug" {
		g.options.HandleWrappers = g.options.HandleWrappers[1:]
	}
	if g.listener == nil {
		listen, err := net.Listen("tcp", g.options.Address)
		if err != nil {
			log.Logger.Panic(err)
		}
		//m := cmux.New(listen)
		//grpcL := m.MatchWithWriters(
		//	cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"),
		//)
		g.options.Address = listen.Addr().String()
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
	ip, _ := utils.GetIntranetIp()
	endpoint := strings.Replace(g.options.Address, "[::]", ip.String(), 1)

	ser := registry.Service{
		Name:     g.options.Name,
		Version:  g.options.Version,
		Node:     utils.GetUuidV3(reg.GetPrefix()),
		Endpoint: endpoint,
		Namespace: g.options.Namespace,
	}
	for _, server := range servers {
		server(&ser)
	}
	return reg.Registry(&ser)
}

func (g *GrpcServer) Server() *grpc.Server {
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
