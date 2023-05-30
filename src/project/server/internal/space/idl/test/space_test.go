package space

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/log"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/check"
	"shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"
)

type SpaceTestSuite struct {
	suite.Suite
	space pb.SpaceServiceClient
}

func (s *SpaceTestSuite) SetupTest() {
	check.Assert(infrastructure.Init(dig.New()).Invoke(func(config *config.Config, consul consul.Consul) error {
		if conn, err := grpc.NewConn(consul.Address(config.Service.Name)); err != nil {
			return err
		} else {
			s.space = pb.NewSpaceServiceClient(conn)
			return nil
		}
	}))
}
func (s *SpaceTestSuite) TestCreate() {
	req := &pb.SpaceCreateRequest{Name: "hi"}
	res, err := s.space.Create(context.TODO(), req)
	if s.Equal(nil, err, "test failure %v", err) {
		s.Equal(req.Name, res.Data.Name)
	}
}
func (s *SpaceTestSuite) TestList() {
	req := &pb.SpaceListRequest{}
	res, err := s.space.List(context.TODO(), req)
	if s.Equal(nil, err, "test failure %v", err) {
		if res.BaseResp != nil && res.BaseResp.Code > 100000 {
			log.Fatalf("test failure: %v", res.BaseResp)
		}
	}
}
func TestSpaceTestSuite(t *testing.T) { suite.Run(t, new(SpaceTestSuite)) }
