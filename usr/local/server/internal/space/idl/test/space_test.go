package space

import (
	"context"
	"testing"

	"shylinux.com/x/golang-story/src/project/server/infrastructure"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/tests"
	"shylinux.com/x/golang-story/src/project/server/internal/space/idl/pb"
)

type SpaceTestSuite struct {
	*tests.Suite
	ctx   context.Context
	space pb.SpaceServiceClient
	id    int64
}

func (s *SpaceTestSuite) SetupTest() {
	s.space = pb.NewSpaceServiceClient(s.Conn(s.ctx, pb.SpaceService_ServiceDesc.ServiceName))
	if res, err := s.space.Create(s.ctx, &pb.SpaceCreateRequest{Name: "goodlife"}); err != nil {
		panic(err)
	} else {
		s.id = res.Data.SpaceID
	}
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
		res, err := s.space.Create(s.ctx, &pb.SpaceCreateRequest{Name: c.name})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func (s *SpaceTestSuite) TestRemove() {
	cases := []struct {
		ok      bool
		spaceID int64
	}{
		{ok: false, spaceID: 0},
		{ok: true, spaceID: s.id},
		{ok: false, spaceID: -1},
	}
	for i, c := range cases {
		res, err := s.space.Remove(s.ctx, &pb.SpaceRemoveRequest{SpaceID: c.spaceID})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func (s *SpaceTestSuite) TestInfo() {
	cases := []struct {
		ok      bool
		spaceID int64
	}{
		{ok: false, spaceID: 0},
		{ok: true, spaceID: s.id},
		{ok: false, spaceID: -1},
	}
	for i, c := range cases {
		res, err := s.space.Info(s.ctx, &pb.SpaceInfoRequest{SpaceID: c.spaceID})
		s.ConveySo(i, c.ok, c, res, err)
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
		res, err := s.space.List(s.ctx, &pb.SpaceListRequest{Page: c.page, Count: c.count})
		s.ConveySo(i, c.ok, c, res, err)
	}
}
func TestSpaceTestSuite(t *testing.T) {
	infrastructure.Test(t, func(suite *tests.Suite) interface{} {
		return &SpaceTestSuite{Suite: suite, ctx: suite.Context()}
	})
}
