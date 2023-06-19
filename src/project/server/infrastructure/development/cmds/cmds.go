package cmds

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/reflect"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

type Cmds struct{ *cobra.Command }

func New(config *config.Config) *Cmds {
	return &Cmds{&cobra.Command{
		Use:   "matrix",
		Short: "matrix devops command",
		Long: `matrix devops command
command
  project create
  service create
  product create
  deploy enviroment
  server controll
  proto define
service
  auth client
  user client
  mesh client
runtime
  java language
  node language
`,
		Run: func(cmd *cobra.Command, arg []string) {
			fmt.Println(system.MarshalIndent(config))
		},
	}}
}
func (s *Cmds) Run() error { return s.Execute() }
func (s *Cmds) Add(name string, info string, cb func(ctx context.Context, arg ...string)) *Cmds {
	cmd := &cobra.Command{
		Use:   name,
		Short: strings.Split(info, "\n")[0],
		Long:  info,
		Run:   func(cmd *cobra.Command, arg []string) { cb(cmd.Context(), arg...) },
	}
	s.AddCommand(cmd)
	return &Cmds{cmd}
}
func (s *Cmds) Register(name string, help string, obj interface{}) *Cmds {
	cmds := s.Add(name, help, func(ctx context.Context, arg ...string) {})
	reflect.MethodList(obj, func(name string, method reflect.Method) {
		name, help = strings.ToLower(name), strings.ToLower(name)
		reflect.FieldList(method.NewParam(1), func(name string, field reflect.Field) {
			help += " " + name
		})
		cmds.Add(name, help, func(ctx context.Context, arg ...string) {
			method.Call(ctx, method.NewParam(1, arg...))
		})
	})
	return cmds
}
