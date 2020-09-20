package selector

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/log"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/serrors"
	"github.com/liujunren93/share/utils"
	"google.golang.org/grpc"
	"time"
)

type (
	Next      func() (*grpc.ClientConn, error)
	RoundType func(s *[]*registry.Service, ctx context.Context, option ...grpc.DialOption) Next
)

func Round(s *[]*registry.Service, ctx context.Context, option ...grpc.DialOption) Next {
	value := ctx.Value("serverName")
	return func() (*grpc.ClientConn, error) {

		if len(*s)==0 {
			err := fmt.Errorf("[share]service:%v is down", value)
			log.Logger.Error(err)
			return nil, serrors.InternalServerError(err)
		}
		randInt := utils.RandInt(len(*s))
		endpoint := (*s)[randInt].Endpoint
		ctx,_:=context.WithTimeout(ctx, time.Second)
		return getConn(endpoint, ctx, option...)
	}
}

func RoundRobin(s *[]*registry.Service, ctx context.Context, option ...grpc.DialOption) Next {
	randInt := utils.RandInt(len(*s))

	return func() (*grpc.ClientConn, error) {
		randInt++
		endpoint := (*s)[randInt%len(*s)].Endpoint
		ctx,_:=context.WithTimeout(ctx, time.Second)

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
