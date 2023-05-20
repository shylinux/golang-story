package test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"shylinux.com/x/golang-story/src/project/server/idl/pb"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
)

type UserTestSuite struct {
	suite.Suite
	addr string
	conn *grpc.ClientConn
	user pb.UserServiceClient
}

func (s *UserTestSuite) SetupTest() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config load failure: %v", err)
	}
	conn, err := grpc.Dial(config.Service.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	s.conn = conn
	s.user = pb.NewUserServiceClient(conn)
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
			log.Fatalf("test failure: %v", res.BaseResp)
		}
	}
}
func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
