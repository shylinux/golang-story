package controller

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
)

type HealthController struct {
	pb.UnimplementedUserServiceServer
}

func NewHealthController() *HealthController {
	return &HealthController{}
}
func (h *HealthController) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}
func (h *HealthController) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}
