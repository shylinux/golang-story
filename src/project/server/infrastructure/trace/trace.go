package trace

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Outgoing(ctx context.Context) map[string]string {
	md, _ := metadata.FromIncomingContext(ctx)
	meta := map[string]string{}
	for k, v := range md {
		meta[k] = v[0]
	}
	return meta
}
func Incoming(ctx context.Context, meta map[string]string) context.Context {
	kv := []string{}
	for k, v := range meta {
		kv = append(kv, k, v)
	}
	return metadata.NewIncomingContext(ctx, metadata.Pairs(kv...))
}
func ServerAccess(ctx context.Context, cb func(context.Context)) {
	otelgrpc.UnaryServerInterceptor()(ctx, nil, &grpc.UnaryServerInfo{}, func(ctx context.Context, req interface{}) (interface{}, error) {
		cb(ctx)
		return nil, nil
	})
}
