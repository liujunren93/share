package wrapper

import "google.golang.org/grpc"

type CallWrapper func() (interceptor grpc.UnaryClientInterceptor, wrapName string)
