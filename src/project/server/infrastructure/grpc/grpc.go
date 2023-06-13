package grpc

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
)

func NewServer(config *config.Config) *grpc.Server {
	defer tracer()()
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(otelgrpc.UnaryServerInterceptor(), serverInterceptor))
	grpc_health_v1.RegisterHealthServer(server, &HealthController{})
	return server
}

type ClientConn struct {
	*grpc.ClientConn
}

func NewConn(ctx context.Context, target string) (*ClientConn, error) {
	defer tracer()()
	conn, err := grpc.DialContext(ctx, target, grpc.WithChainUnaryInterceptor(otelgrpc.UnaryClientInterceptor(), clientInterceptor),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`), grpc.WithInsecure())
	return &ClientConn{conn}, err
}
