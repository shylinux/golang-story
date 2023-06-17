package api

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
	"shylinux.com/x/golang-story/src/project/server/internal/mesh/idl/pb"
)

func NewMachineServiceClient(ctx context.Context, consul consul.Consul) (pb.MachineServiceClient, error) {
	if conn, err := grpc.NewConn(ctx, consul.Address(pb.MachineService_ServiceDesc.ServiceName)); err != nil {
		return nil, err
	} else {
		return pb.NewMachineServiceClient(conn), err
	}
}
