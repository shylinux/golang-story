package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/check"
)

type UserTestSuite struct {
	suite.Suite
	user pb.UserServiceClient
}

func (s *UserTestSuite) SetupTest() {
	check.Assert(infrastructure.Init(dig.New()).Invoke(func(config *config.Config, consul consul.Consul) error {
		if conn, err := grpc.NewConn(context.TODO(), consul.Address(config.Service.Name)); err != nil {
			return err
		} else {
			s.user = pb.NewUserServiceClient(conn)
			return nil
		}
	}))
}
func (s *UserTestSuite) TestCreate() {
	req := &pb.UserCreateRequest{Name: "hi"}
	res, err := s.user.Create(context.TODO(), req)
	if s.Equal(nil, err, "test failure %v", err) {
		s.Equal(req.Name, res.Data.Name)
	}
}
func (s *UserTestSuite) TestList() {
	req := &pb.UserListRequest{}
	res, err := s.user.List(context.TODO(), req)
	if s.Equal(nil, err, "test failure %v", err) {
		if res.BaseResp != nil && res.BaseResp.Code > 100000 {
			logs.Fatalf("test failure: %v", res.BaseResp)
		}
	}
}
func TestUserTestSuite(t *testing.T) { suite.Run(t, new(UserTestSuite)) }
