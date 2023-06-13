package api

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
)

func NewUserServiceClient(ctx context.Context, consul consul.Consul) (pb.UserServiceClient, error) {
	if conn, err := grpc.NewConn(ctx, consul.Address(pb.UserService_ServiceDesc.ServiceName)); err != nil {
		return nil, err
	} else {
		return pb.NewUserServiceClient(conn), err
	}
}
