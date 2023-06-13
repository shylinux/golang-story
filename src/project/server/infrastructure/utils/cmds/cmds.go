package cmds

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

type Cmds struct {
	*cobra.Command
}

func New(config *config.Config) *Cmds {
	return &Cmds{&cobra.Command{
		Use:   config.Server.Name,
		Short: config.Server.Name + " command",
		Long:  config.Server.Name + " command",
		Run: func(cmd *cobra.Command, arg []string) {
			fmt.Println(logs.MarshalIndent(config))
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
	t, v := reflect.TypeOf(obj), reflect.ValueOf(obj)
	for i := 0; i < v.NumMethod(); i++ {
		method, name := v.Method(i), strings.ToLower(t.Method(i).Name)
		cmds.Add(name, name, func(ctx context.Context, arg ...string) {
			method.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(Bind(reflect.New(method.Type().In(1).Elem()).Interface(), arg...))})
		})
	}
	return cmds
}
func Bind(req interface{}, arg ...string) interface{} {
	rt, rv := reflect.TypeOf(req).Elem(), reflect.ValueOf(req).Elem()
	trans := map[string]string{}
	for i := 0; i < rv.NumField(); i++ {
		if unicode.IsUpper(rune(rt.Field(i).Name[0])) {
			trans[strings.ToLower(rt.Field(i).Name)] = rt.Field(i).Name
		}
	}
	for i := 0; i < len(arg); i += 2 {
		if fv := rv.FieldByName(trans[arg[i]]); fv.CanSet() {
			switch fv.Type().Kind() {
			case reflect.String:
				fv.SetString(arg[i+1])
			case reflect.Int64:
				v, _ := strconv.ParseInt(arg[i+1], 10, 64)
				fv.SetInt(v)
			}
		}
	}
	return req
}
