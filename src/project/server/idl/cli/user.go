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

type UserServiceCmds struct {
	consul consul.Consul
	client pb.UserServiceClient
}

func NewUserServiceCmds(cmds *cmds.Cmds, consul consul.Consul) (*UserServiceCmds, error) {
	_cmds := &UserServiceCmds{consul: consul}
	cmds.Register("user", "user", _cmds)
	return _cmds, nil
}

func (s *UserServiceCmds) conn(ctx context.Context, arg ...string) {
	if s.client != nil {
		return
	}
	conn, err := grpc.NewConn(ctx, s.consul.Address(pb.UserService_ServiceDesc.ServiceName))
	if err != nil {
		return
	}
	s.client = pb.NewUserServiceClient(conn)
}

func (s *UserServiceCmds) Create(ctx context.Context, req *pb.UserCreateRequest) {
	s.conn(ctx)
	if res, err := s.client.Create(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *UserServiceCmds) Remove(ctx context.Context, req *pb.UserRemoveRequest) {
	s.conn(ctx)
	if res, err := s.client.Remove(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *UserServiceCmds) Rename(ctx context.Context, req *pb.UserRenameRequest) {
	s.conn(ctx)
	if res, err := s.client.Rename(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *UserServiceCmds) Search(ctx context.Context, req *pb.UserSearchRequest) {
	s.conn(ctx)
	if res, err := s.client.Search(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *UserServiceCmds) Info(ctx context.Context, req *pb.UserInfoRequest) {
	s.conn(ctx)
	if res, err := s.client.Info(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}

func (s *UserServiceCmds) List(ctx context.Context, req *pb.UserListRequest) {
	s.conn(ctx)
	if res, err := s.client.List(ctx, req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", logs.MarshalIndent(res))
	}
}
