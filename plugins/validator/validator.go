package validator

import (
	"context"
	"github.com/liujunren93/share/errors"
	"google.golang.org/grpc"
)

type Validator interface {
	Validate() error
}

func NewHandlerWrapper() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if v, ok := req.(Validator); ok {
			if err := v.Validate(); err != nil {

				return nil, errors.BadRequest(err, nil)
			}
		}
		return handler(ctx, req)
	}

}
