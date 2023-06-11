package proxy

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	grpcs "shylinux.com/x/golang-story/src/project/server/infrastructure/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/trace"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/check"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/metadata"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/response"
)

type Proxy struct {
	Auth  func(context.Context, string, string) (context.Context, error)
	proxy map[string]reflect.Value
	*config.Config
	consul.Consul
}

func New(config *config.Config, consul consul.Consul) *Proxy {
	return &Proxy{Auth: func(ctx context.Context, api string, token string) (context.Context, error) {
		return ctx, nil
	}, proxy: map[string]reflect.Value{}, Config: config, Consul: consul}
}
func (s *Proxy) Register(service string, controller interface{}) {
	t, v := reflect.TypeOf(controller), reflect.ValueOf(controller)
	for i := 0; i < v.NumMethod(); i++ {
		if unicode.IsLower(rune(t.Method(i).Name[0])) || v.Method(i).Type().NumIn() != 2 || v.Method(i).Type().NumOut() != 2 {
			continue
		}
		logs.Infof("proxy register %s/%s", service, t.Method(i).Name)
		s.proxy[path.Join("/", service, t.Method(i).Name)] = v.Method(i)
	}
}
func (s *Proxy) handler(ctx *gin.Context) {
	trace.ServerAccess(context.Background(), func(_ctx context.Context) {
		var err error
		var res interface{}
		ctx.Writer.Header().Add("TraceID", logs.TraceID(_ctx))
		begin, p := time.Now(), strings.TrimPrefix(ctx.Request.URL.Path, "/api")
		echo := func(res interface{}, err error) {
			if err != nil && err.Error() != "" {
				logs.Warnf("proxy result %s %s cost:%s", p, err.Error(), logs.Cost(begin), _ctx)
			} else {
				logs.Infof("proxy result %s %s cost:%s", p, logs.Marshal(res), logs.Cost(begin), _ctx)
			}
			response.WriteData(ctx, res, err)
		}
		if _ctx, err = s.Auth(_ctx, p, strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")); err != nil {
			echo(nil, errors.NewNotAuth(err))
		} else if cb, ok := s.proxy[p]; !ok {
			echo(nil, errors.NewNotFoundProxy(fmt.Errorf(p)))
		} else if arg, err := s.parse(ctx, cb); err != nil {
			echo(nil, errors.NewInvalidParams(err))
		} else if logs.Infof("proxy access %s %s username:%s %s %s", p, logs.Marshal(arg), metadata.GetValue(_ctx, metadata.USERNAME), ctx.RemoteIP(), ctx.GetHeader("User-Agent"), _ctx); s.Config.Proxy.Local {
			res := cb.Call([]reflect.Value{reflect.ValueOf(_ctx), reflect.ValueOf(arg)})
			echo(res[0].Interface(), errors.ParseResp(res[1].Interface().(error), p))
		} else {
			conn, err := grpcs.NewConn(_ctx, s.Consul.Address(strings.Split(p, "/")[1]))
			if res = reflect.New(cb.Type().Out(0).Elem()).Interface(); err == nil {
				err = conn.Invoke(_ctx, p, arg, res)
			}
			echo(res, errors.ParseResp(err, p))
		}
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
			k = strings.ToLower(k)
			for _, val := range val {
				if field, ok := t.FieldByName(trans[k]); ok {
					switch field.Type.Kind() {
					case reflect.String:
						v.FieldByName(trans[k]).SetString(val)
					case reflect.Int64:
						i, _ := strconv.ParseInt(val, 10, 64)
						v.FieldByName(trans[k]).SetInt(i)
					}
				}
			}
		}
	} else if err := ctx.Bind(arg); err != nil {
		return nil, err
	}
	if valid, ok := arg.(interface{ Validate() error }); ok {
		if err := valid.Validate(); err != nil {
			return nil, err
		}
	}
	return arg, nil
}
func (s *Proxy) Run(conf config.Proxy) {
	if !path.IsAbs(conf.Root) {
		wd, _ := os.Getwd()
		conf.Root = path.Join(wd, conf.Root)
	}
	engine := gin.New()
	engine.POST("/*s", s.handler)
	engine.GET("/*s", func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/api/") {
			s.handler(ctx)
		} else if ctx.Request.URL.Path == "/" {
			ctx.File(path.Join(conf.Root, "index.html"))
		} else {
			ctx.File(path.Join(conf.Root, ctx.Request.URL.Path))
		}
	})
	logs.Infof("proxy start %s:%d root %s", conf.Host, conf.Port, conf.Root)
	check.Assert(engine.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port)))
}
