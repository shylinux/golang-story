package trace

import (
	"context"
	"fmt"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func TraceID(ctx context.Context) string {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return fmt.Sprintf("%s-%s", span.TraceID().String(), span.SpanID())
	}
	return ""
}
func ServerAccess(ctx context.Context, cb func(context.Context)) {
	otelgrpc.UnaryServerInterceptor()(ctx, nil, &grpc.UnaryServerInfo{}, func(ctx context.Context, req interface{}) (interface{}, error) {
		cb(ctx)
		return nil, nil
	})
}
