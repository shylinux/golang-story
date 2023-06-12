package cmds

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

type Cmds struct {
	*cobra.Command
}

func New() *Cmds {
	return &Cmds{&cobra.Command{
		Use:   "demo",
		Short: "demo command",
		Long:  "demo command",
		Run: func(cmd *cobra.Command, arg []string) {
			fmt.Println(logs.FileLine(1))
		},
	}}
}
func (s *Cmds) Run() error {
	return s.Execute()
}
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
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumMethod(); i++ {
		name := strings.ToLower(t.Method(i).Name)
		method := v.Method(i)
		cmds.Add(name, name, func(ctx context.Context, arg ...string) {
			args := []reflect.Value{reflect.ValueOf(ctx)}
			method.Call(args)
		})
	}
	return cmds
}
