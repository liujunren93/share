package recover

import (
	"context"
	"github.com/liujunren93/share/log"
	"google.golang.org/grpc"
)

type Recover struct {

}

func NewHandlerWrapper() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if err:=recover();err!=nil {
				log.Logger.Error(err)
			}
		}()
		return handler(ctx, req)
	}
}