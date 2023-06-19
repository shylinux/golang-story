package api

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
)

func NewServiceServiceClient(ctx context.Context, consul consul.Consul) (pb.ServiceServiceClient, error) {
	if conn, err := grpc.NewConn(ctx, consul.Address(pb.ServiceService_ServiceDesc.ServiceName)); err != nil {
		return nil, err
	} else {
		return pb.NewServiceServiceClient(conn), err
	}
}
