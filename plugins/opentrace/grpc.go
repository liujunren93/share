package opentrace

import (
	"context"
	log2 "github.com/liujunren93/share/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"

	"google.golang.org/grpc/metadata"
)

type MDCarrier struct {
	metadata.MD
}

func (m MDCarrier) ForeachKey(handler func(key, val string) error) error {
	for k, strs := range m.MD {
		for _, v := range strs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}
func (m MDCarrier) Set(key, val string) {
	m.MD[key] = append(m.MD[key], val)
}

func ServerGrpcWrap(ot opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		spanContext, err := ot.Extract(
			opentracing.TextMap,
			MDCarrier{md},
		)
		var serSpan opentracing.Span
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			log2.Logger.Errorf("extract from metadata err: %v", err)
		} else {
			serSpan := ot.StartSpan(
				info.FullMethod,
				ext.RPCServerOption(spanContext),
				opentracing.Tag{Key: string(ext.Component), Value: "gRPC Server"},
				ext.SpanKindRPCServer,
			)
			defer serSpan.Finish()

			ctx = opentracing.ContextWithSpan(ctx, serSpan)
		}

		i, err := handler(ctx, req)
		if err != nil {
			log2.Logger.Error(err)
			serSpan.LogFields(log.String("event", "error"), log.String("message", err.Error()))
		}
		return i, err
	}
}

func ClientGrpcCallWrap(ot opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		//一个RPC调用的服务端的span，和RPC服务客户端的span构成ChildOf关系
		var parentCtx opentracing.SpanContext
		parentSpan := opentracing.SpanFromContext(ctx)
		if parentSpan != nil {
			parentCtx = parentSpan.Context()
		}

		span := ot.StartSpan(
			method,
			opentracing.ChildOf(parentCtx),
			opentracing.Tag{Key: string(ext.Component), Value: "gRPC Client"},
			ext.SpanKindRPCClient,
		)
		defer span.Finish()
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		err := ot.Inject(
			span.Context(),
			opentracing.TextMap,
			MDCarrier{md}, // 自定义 carrier
		)
		if err != nil {
			log.Error(err)
		}
		newCtx := metadata.NewOutgoingContext(ctx, md)
		err = invoker(newCtx, method, req, reply, cc, opts...)

		if err != nil {
			log2.Logger.Error(err)
			span.LogFields(log.String("event", "error"), log.String("message", err.Error()))
		}

		return err
	}
}
