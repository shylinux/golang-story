package cli

import (
	"context"
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
)

type ServiceServiceCmds struct {
	consul consul.Consul
	client pb.ServiceServiceClient
}

func NewServiceServiceCmds(cmds *cmds.Cmds, consul consul.Consul) (*ServiceServiceCmds, error) {
	_cmds := &ServiceServiceCmds{consul: consul}
	cmds.Register("service", "service service client", _cmds)
	return _cmds, nil
}

func (s *ServiceServiceCmds) conn(ctx context.Context, arg ...string) {
	if s.client != nil {
		return
	}
	conn, err := grpc.NewConn(ctx, s.consul.Address(pb.ServiceService_ServiceDesc.ServiceName))
	if err != nil {
		return
	}
	s.client = pb.NewServiceServiceClient(conn)
}

func (s *ServiceServiceCmds) Create(ctx context.Context, req *pb.ServiceCreateRequest) {
	s.conn(ctx)
	if res, err := s.client.Create(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *ServiceServiceCmds) Remove(ctx context.Context, req *pb.ServiceRemoveRequest) {
	s.conn(ctx)
	if res, err := s.client.Remove(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *ServiceServiceCmds) Inputs(ctx context.Context, req *pb.ServiceInputsRequest) {
	s.conn(ctx)
	if res, err := s.client.Inputs(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *ServiceServiceCmds) Deploy(ctx context.Context, req *pb.ServiceDeployRequest) {
	s.conn(ctx)
	if res, err := s.client.Deploy(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *ServiceServiceCmds) Info(ctx context.Context, req *pb.ServiceInfoRequest) {
	s.conn(ctx)
	if res, err := s.client.Info(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *ServiceServiceCmds) List(ctx context.Context, req *pb.ServiceListRequest) {
	s.conn(ctx)
	if res, err := s.client.List(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}
