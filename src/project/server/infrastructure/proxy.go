package infrastructure

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/response"
)

type Proxy struct {
	proxy map[string]reflect.Value
}

func NewProxy() *Proxy {
	return &Proxy{map[string]reflect.Value{}}
}
func (s *Proxy) Register(service string, controller interface{}) {
	t := reflect.TypeOf(controller)
	v := reflect.ValueOf(controller)
	for i := 0; i < v.NumMethod(); i++ {
		logs.Infof("register proxy %s/%s", service, t.Method(i).Name)
		s.proxy[path.Join("/", service, t.Method(i).Name)] = v.Method(i)
	}
}
func (s *Proxy) handler(ctx *gin.Context) {
	p := strings.TrimPrefix(ctx.Request.URL.Path, "/api")
	cb, ok := s.proxy[p]
	if !ok {
		response.WriteError(ctx, errors.NewNotFoundProxy(fmt.Errorf(p)))
		return
	}
	arg, err := s.parse(ctx, cb)
	if err != nil {
		response.WriteError(ctx, errors.NewInvalidParams(err))
		return
	}
	otelgrpc.UnaryServerInterceptor()(context.Background(), nil, &grpc.UnaryServerInfo{}, func(_ctx context.Context, req interface{}) (interface{}, error) {
		logs.Infof("proxy %s %+v", p, arg, _ctx)
		res := cb.Call([]reflect.Value{reflect.ValueOf(_ctx), reflect.ValueOf(arg)})
		if err, ok := res[1].Interface().(error); ok && err != nil {
			logs.Warnf("proxy %s %s", p, err.Error(), _ctx)
			if ls := strings.SplitN(err.Error(), ":", 2); len(ls) > 1 {
				if code, _err := strconv.ParseInt(ls[0], 10, 64); _err == nil && code > 0 {
					response.WriteError(ctx, errors.NewResp(err, code, ls[1]))
					return nil, nil
				}
			}
			response.WriteError(ctx, errors.NewResp(err, enums.Errors.Unknown, p))
		} else {
			response.WriteData(ctx, res[0].Interface(), nil)
		}
		return nil, nil
	})

}
func (s *Proxy) parse(ctx *gin.Context, cb reflect.Value) (interface{}, error) {
	t := cb.Type()
	arg := reflect.New(t.In(1).Elem()).Interface()
	if ctx.Request.Method == http.MethodGet {
		t := reflect.TypeOf(arg).Elem()
		v := reflect.ValueOf(arg).Elem()
		trans := map[string]string{}
		for i := 0; i < t.NumField(); i++ {
			name := t.Field(i).Name
			if key := strings.ToLower(name); key != name {
				trans[key] = name
			}
		}
		ctx.Request.ParseForm()
		for k, val := range ctx.Request.Form {
			for _, val := range val {
				if field, ok := t.FieldByName(trans[k]); ok {
					switch field.Type.Kind() {
					case reflect.Int64:
						i, _ := strconv.ParseInt(val, 10, 64)
						v.FieldByName(trans[k]).SetInt(i)
					case reflect.String:
						v.FieldByName(trans[k]).SetString(val)
					}
				}
			}
		}
	} else if err := ctx.Bind(arg); err != nil {
		logs.Infof("%s %s %+v", ctx.Request.Method, ctx.Request.URL, err)
		ctx.JSON(http.StatusInternalServerError, "")
		return nil, errors.New(fmt.Errorf("error"), "")
	}
	return arg, nil
}
func (s *Proxy) run(conf config.Gateway) {
	engine := gin.New()
	engine.GET("/api/*s", s.handler)
	engine.POST("/api/*s", s.handler)
	if !path.IsAbs(conf.Root) {
		wd, _ := os.Getwd()
		conf.Root = path.Join(wd, conf.Root)
	}
	engine.Static("/js", path.Join(conf.Root, "js"))
	engine.Static("/css", path.Join(conf.Root, "css"))
	engine.StaticFile("/", path.Join(conf.Root, "index.html"))
	engine.StaticFile("/favicon.ico", path.Join(conf.Root, "favicon.ico"))
	logs.Infof("gateway %s", conf.Root)
	if err := engine.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
		panic(errors.New(err, "start gin failure"))
	}
}
