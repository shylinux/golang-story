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

type MachineServiceCmds struct {
	consul consul.Consul
	client pb.MachineServiceClient
}

func NewMachineServiceCmds(cmds *cmds.Cmds, consul consul.Consul) (*MachineServiceCmds, error) {
	_cmds := &MachineServiceCmds{consul: consul}
	cmds.Register("machine", "machine service client", _cmds)
	return _cmds, nil
}

func (s *MachineServiceCmds) conn(ctx context.Context, arg ...string) {
	if s.client != nil {
		return
	}
	conn, err := grpc.NewConn(ctx, s.consul.Address(pb.MachineService_ServiceDesc.ServiceName))
	if err != nil {
		return
	}
	s.client = pb.NewMachineServiceClient(conn)
}

func (s *MachineServiceCmds) Create(ctx context.Context, req *pb.MachineCreateRequest) {
	s.conn(ctx)
	if res, err := s.client.Create(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *MachineServiceCmds) Remove(ctx context.Context, req *pb.MachineRemoveRequest) {
	s.conn(ctx)
	if res, err := s.client.Remove(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *MachineServiceCmds) Change(ctx context.Context, req *pb.MachineChangeRequest) {
	s.conn(ctx)
	if res, err := s.client.Change(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *MachineServiceCmds) Info(ctx context.Context, req *pb.MachineInfoRequest) {
	s.conn(ctx)
	if res, err := s.client.Info(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *MachineServiceCmds) List(ctx context.Context, req *pb.MachineListRequest) {
	s.conn(ctx)
	if res, err := s.client.List(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}
