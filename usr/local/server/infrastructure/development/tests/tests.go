package tests

import (
	"context"
	"fmt"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

type Suite struct {
	suite.Suite
	consul.Consul
	ctx context.Context
}

func New(config *config.Config, logger logs.Logger, consul consul.Consul, ctx context.Context) *Suite {
	return &Suite{Consul: consul, ctx: ctx}
}
func (s *Suite) Context() context.Context {
	return s.ctx
}
func (s *Suite) Run(t *testing.T, ts interface{}) {
	suite.Run(t, ts.(suite.TestingSuite))
}
func (s *Suite) Load(file string, data interface{}) error {
	if f, e := system.Open(file); e != nil {
		return e
	} else {
		defer f.Close()
		return yaml.NewDecoder(f).Decode(data)
	}
}
func (s *Suite) Conn(ctx context.Context, name string) *grpc.ClientConn {
	if conn, err := grpc.NewConn(ctx, s.Consul.Address(name)); err != nil {
		panic(err)
	} else {
		return conn
	}
}
func (s *Suite) ConveySo(i int, ok bool, arg interface{}, res interface{}, err error) {
	Convey(fmt.Sprintf("%s case: %d %+v %v %v", strings.TrimPrefix(errors.FuncName(2), "command-line-arguments."), i+1, arg, logs.Marshal(res), err), s.T(), func() {
		So(ok && err != nil || !ok && err == nil, ShouldBeFalse)
	})
}
