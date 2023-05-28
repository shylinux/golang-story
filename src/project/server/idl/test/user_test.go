package test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"testing"
	"time"

	"github.com/hashicorp/consul/api"
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
	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	conf := api.DefaultConfig()
	conf.Address = config.Consul.Addr
	registry, err := api.NewClient(conf)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := grpc.Dial(config.Service.Name, grpc.WithInsecure(), grpc.WithDialer(func(name string, timeout time.Duration) (net.Conn, error) {
		list, _, err := registry.Health().Service(name, "", true, nil)
		if err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, fmt.Errorf("error")
		}
		i := rand.Intn(len(list))
		log.Printf("what %d %d %#v", i, len(list), list[i].Service.Address)
		return net.Dial("tcp", list[i].Service.Address)
	}))
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
