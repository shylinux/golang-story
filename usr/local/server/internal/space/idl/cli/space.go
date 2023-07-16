package cli

import (
	"context"
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
	"shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"
)

type SpaceServiceCmds struct {
	consul consul.Consul
	client pb.SpaceServiceClient
}

func NewSpaceServiceCmds(cmds *cmds.Cmds, consul consul.Consul) (*SpaceServiceCmds, error) {
	_cmds := &SpaceServiceCmds{consul: consul}
	cmds.Register("space", "space service client", _cmds)
	return _cmds, nil
}

func (s *SpaceServiceCmds) conn(ctx context.Context, arg ...string) {
	if s.client != nil {
		return
	}
	conn, err := grpc.NewConn(ctx, s.consul.Address(pb.SpaceService_ServiceDesc.ServiceName))
	if err != nil {
		return
	}
	s.client = pb.NewSpaceServiceClient(conn)
}

func (s *SpaceServiceCmds) Create(ctx context.Context, req *pb.SpaceCreateRequest) {
	s.conn(ctx)
	if res, err := s.client.Create(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *SpaceServiceCmds) Remove(ctx context.Context, req *pb.SpaceRemoveRequest) {
	s.conn(ctx)
	if res, err := s.client.Remove(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *SpaceServiceCmds) Info(ctx context.Context, req *pb.SpaceInfoRequest) {
	s.conn(ctx)
	if res, err := s.client.Info(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *SpaceServiceCmds) List(ctx context.Context, req *pb.SpaceListRequest) {
	s.conn(ctx)
	if res, err := s.client.List(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}
