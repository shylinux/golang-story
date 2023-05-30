package grpc

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
)

func NewServer(config *config.Config, csl consul.Consul) *grpc.Server {
	// Set up OTLP tracing (stdout for debug).
	exporter, _ := stdout.New(stdout.WithPrettyPrint())
	otel.SetTracerProvider(sdktrace.NewTracerProvider(sdktrace.WithSampler(sdktrace.AlwaysSample()), sdktrace.WithBatcher(exporter)))
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	defer func() { _ = exporter.Shutdown(context.Background()) }()

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		otelgrpc.UnaryServerInterceptor(),
		serverInterceptor,
	))
	grpc_health_v1.RegisterHealthServer(server, &HealthController{})
	return server
}
func NewConn(ctx context.Context, target string) (*grpc.ClientConn, error) {
	// Set up OTLP tracing (stdout for debug).
	exporter, _ := stdout.New(stdout.WithPrettyPrint())
	otel.SetTracerProvider(sdktrace.NewTracerProvider(sdktrace.WithSampler(sdktrace.AlwaysSample()), sdktrace.WithBatcher(exporter)))
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	defer func() { _ = exporter.Shutdown(context.Background()) }()

	return grpc.DialContext(ctx, target, grpc.WithChainUnaryInterceptor(
		otelgrpc.UnaryClientInterceptor(),
		clientInterceptor,
	), grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`), grpc.WithInsecure())
}
