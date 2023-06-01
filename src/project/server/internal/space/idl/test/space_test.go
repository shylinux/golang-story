package space

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/check"
	"shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"
)

type SpaceTestSuite struct {
	suite.Suite
	user pb.SpaceServiceClient
	ctx  context.Context
	t    *testing.T
}

func (s *SpaceTestSuite) SetupTest() {
	check.Assert(infrastructure.Init(dig.New()).Invoke(func(config *config.Config, consul consul.Consul) error {
		if conn, err := grpc.NewConn(s.ctx, consul.Address(pb.SpaceService_ServiceDesc.ServiceName)); err != nil {
			return err
		} else {
			s.user = pb.NewSpaceServiceClient(conn)
			return nil
		}
	}))
}
func (s *SpaceTestSuite) TestCreate() {
	cases := []struct {
		ok   bool
		name string
	}{
		{ok: false, name: ""},
		{ok: false, name: "hi"},
		{ok: true, name: "goodlife"},
	}
	for i, c := range cases {
		_, err := s.user.Create(s.ctx, &pb.SpaceCreateRequest{Name: c.name})
		Convey(fmt.Sprintf("%s case: %d %+v", logs.FuncName(1), i+1, c), s.t, func() {
			So(c.ok && err != nil || !c.ok && err == nil, ShouldBeFalse)
		})
	}
}
func (s *SpaceTestSuite) TestInfo() {
	cases := []struct {
		ok bool
		id int64
	}{
		{ok: false, id: 0},
		{ok: true, id: 1},
		{ok: false, id: -1},
	}
	for i, c := range cases {
		_, err := s.user.Info(s.ctx, &pb.SpaceInfoRequest{Id: c.id})
		Convey(fmt.Sprintf("%s case: %d %+v", logs.FuncName(1), i+1, c), s.t, func() {
			So(c.ok && err != nil || !c.ok && err == nil, ShouldBeFalse)
		})
	}
}
func (s *SpaceTestSuite) TestList() {
	cases := []struct {
		ok    bool
		page  int64
		count int64
	}{
		{ok: false, page: 0, count: 0},
		{ok: false, page: 0, count: 10},
		{ok: true, page: 1, count: 10},
		{ok: true, page: 1, count: 10},
	}
	for i, c := range cases {
		_, err := s.user.List(s.ctx, &pb.SpaceListRequest{Page: c.page, Count: c.count})
		Convey(fmt.Sprintf("%s case: %d %+v", logs.FuncName(1), i+1, c), s.t, func() {
			So(c.ok && err != nil || !c.ok && err == nil, ShouldBeFalse)
		})
	}
}
func TestSpaceTestSuite(t *testing.T) { suite.Run(t, &SpaceTestSuite{ctx: context.TODO(), t: t}) }
