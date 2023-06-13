package api

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"
)

func NewSpaceServiceClient(ctx context.Context, consul consul.Consul) (pb.SpaceServiceClient, error) {
	if conn, err := grpc.NewConn(ctx, consul.Address(pb.SpaceService_ServiceDesc.ServiceName)); err != nil {
		return nil, err
	} else {
		return pb.NewSpaceServiceClient(conn), err
	}
}
