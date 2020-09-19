package selector

import (
	"context"
	"github.com/liujunren93/share/log"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/serrors"
	"github.com/liujunren93/share/utils"
	"google.golang.org/grpc"
)

type (
	Next      func() (*grpc.ClientConn, error)
	RoundType func(s *[]*registry.Service, ctx context.Context, option ...grpc.DialOption) Next
)

func Round(s *[]*registry.Service, ctx context.Context, option ...grpc.DialOption) Next {
	return func() (*grpc.ClientConn, error) {
		randInt := utils.RandInt(len(*s))
		endpoint := (*s)[randInt].Endpoint
		return getConn(endpoint, ctx, option...)
	}
}

func roundRobin(s *[]*registry.Service, ctx context.Context, option ...grpc.DialOption) Next {
	randInt := utils.RandInt(len(*s))
	return func() (*grpc.ClientConn, error) {
		endpoint := (*s)[randInt%randInt].Endpoint
		return getConn(endpoint, ctx, option...)
	}
}

func getConn(endpoint string, ctx context.Context, option ...grpc.DialOption) (*grpc.ClientConn, error) {
	dialContext, err := grpc.DialContext(ctx, endpoint, option...)
	if err != nil {
		log.Logger.Errorf("[share]service:%v", err)
		return nil, serrors.InternalServerError(err)
	}
	return dialContext, nil
}
