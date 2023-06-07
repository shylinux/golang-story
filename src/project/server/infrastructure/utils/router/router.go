package router

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/response"
)

func init() { os.RemoveAll("api") }

func saveapi(grp *gin.RouterGroup, name string, method interface{}) {
	if t := reflect.TypeOf(method); t.NumIn() == 2 {
		arg := t.In(1).Elem()
		args := ""
		args += "{"
		for i := 0; i < arg.NumField(); i++ {
			args += fmt.Sprintf(`"%s":"",`, strings.ToLower(arg.Field(i).Name))
		}
		args = strings.TrimSuffix(args, ",")
		args += "}"
		p := path.Join("api", grp.BasePath(), name)
		os.MkdirAll(path.Dir(p), 0755)
		ioutil.WriteFile(p, []byte(args), 0644)
	}
}
func Register(g *gin.Engine, group string, controller interface{}) {
	grp := g.Group(group)
	t := reflect.TypeOf(controller)
	v := reflect.ValueOf(controller)
	for i := 0; i < v.NumMethod(); i++ {
		name := strings.ToLower(t.Method(i).Name)
		method := v.Method(i).Interface()
		saveapi(grp, name, method)
		grp.POST(name, handler(method))
	}
}
func handler(method interface{}) func(*gin.Context) {
	return func(ctx *gin.Context) {
		t := reflect.TypeOf(method)
		v := reflect.ValueOf(method)
		var res []reflect.Value
		switch t.NumIn() {
		case 1:
			logs.Infof("%s %s", ctx.Request.Method, ctx.Request.URL)
			res = v.Call([]reflect.Value{reflect.ValueOf(ctx)})
		case 2:
			arg := reflect.New(t.In(1).Elem()).Interface()
			if err := ctx.Bind(arg); err != nil {
				logs.Infof("%s %s %+v", ctx.Request.Method, ctx.Request.URL, err)
				response.WriteError(ctx, err)
				return
			}
			logs.Infof("%s %s %+v", ctx.Request.Method, ctx.Request.URL, arg)
			res = v.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(arg)})
		default:
			response.WriteError(ctx, fmt.Errorf("func arg must be: (ctx, [data])"))
		}
		if len(res) == 0 {
			return
		}
		switch err := res[len(res)-1].Interface().(type) {
		case nil:
			if len(res) == 1 {
				response.WriteError(ctx, nil)
				return
			}
		case error:
			response.WriteError(ctx, err)
			return
		default:
			response.WriteError(ctx, fmt.Errorf("last res must be error %v", err))
			return
		}
		response.WriteData(ctx, res[0].Interface(), res[1].Interface())
	}
}
