package proxy

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/grpc"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/metadata"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/reflect"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/response"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/trace"
)

type Proxy struct {
	Auth  func(context.Context, string, string) (context.Context, error)
	proxy map[string]reflect.Method
	*config.Config
	consul.Consul
}

func New(config *config.Config, consul consul.Consul) *Proxy {
	return &Proxy{
		Auth:   func(ctx context.Context, api string, token string) (context.Context, error) { return ctx, nil },
		proxy:  map[string]reflect.Method{},
		Consul: consul,
		Config: config,
	}
}
func (s *Proxy) Register(service string, controller interface{}) {
	reflect.MethodList(controller, func(name string, method reflect.Method) {
		logs.Infof("proxy register %s/%s", service, name)
		s.proxy[path.Join(service, name)] = method
	})
}
func (s *Proxy) Run() error {
	conf := s.Config.Proxy
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.POST("/*s", s.handler)
	engine.GET("/*s", func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/api/menu/") {
			ctx.JSON(http.StatusOK, s.Config.Product)
		} else if strings.HasPrefix(ctx.Request.URL.Path, "/api/") {
			s.handler(ctx)
		} else if logs.Infof("static %v", ctx.Request.URL.Path); conf.Target != "" {
			if url, err := url.Parse(conf.Target); err != nil {
				logs.Errorf("Target %s %s", conf.Target, err)
			} else {
				httputil.NewSingleHostReverseProxy(url).ServeHTTP(ctx.Writer, ctx.Request)
			}
		} else if ctx.Request.URL.Path == "/" {
			ctx.File(path.Join(conf.Root, "index.html"))
		} else {
			ctx.File(path.Join(conf.Root, ctx.Request.URL.Path))
		}

	})
	addr := config.Address(conf.Host, conf.Port)
	logs.Infof("proxy start %s root %s", addr, conf.Root)
	system.Printfln("proxy start %s", addr)
	errors.Assert(engine.Run(addr))
	return nil
}
func (s *Proxy) handler(ctx *gin.Context) {
	trace.ServerAccess(context.Background(), func(_ctx context.Context) {
		var err error
		var res interface{}
		ctx.Writer.Header().Add("TraceID", trace.TraceID(_ctx))
		begin, api := time.Now(), strings.TrimPrefix(ctx.Request.URL.Path, "/api/")
		echo := func(res interface{}, err error) {
			if err != nil && err.Error() != "" {
				logs.Warnf("proxy result %s %s cost:%s", api, err.Error(), logs.Cost(begin), _ctx)
			} else {
				logs.Infof("proxy result %s %s cost:%s", api, logs.Marshal(res), logs.Cost(begin), _ctx)
			}
			response.WriteData(ctx, res, err)
		}
		if _ctx, err = s.Auth(_ctx, api, strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")); err != nil {
			echo(nil, errors.NewNotAuth(err))
		} else if method, ok := s.proxy[api]; !ok {
			echo(nil, errors.NewNotFoundProxy(fmt.Errorf(api)))
		} else if arg, err := s.parse(ctx, method); err != nil {
			echo(nil, errors.NewInvalidParams(err))
		} else if logs.Infof("proxy access %s %s username:%s %s %s", api, logs.Marshal(arg), metadata.GetValue(_ctx, metadata.USERNAME), ctx.RemoteIP(), ctx.GetHeader("User-Agent"), _ctx); s.Config.Proxy.Local {
			res := method.Call(_ctx, arg)
			echo(res[0], errors.ParseResp(res[1].(error), api))
		} else {
			conn, err := grpc.NewConn(_ctx, s.Consul.Address(strings.Split(api, "/")[0]))
			if err == nil {
				res = method.NewResult(0)
				err = conn.Invoke(_ctx, api, arg, res)
			}
			echo(res, errors.ParseResp(err, api))
		}
	})
}
func (s *Proxy) parse(ctx *gin.Context, method reflect.Method) (interface{}, error) {
	req := method.NewParam(1)
	if ctx.Request.Method == http.MethodGet {
		arg := []string{}
		for k, val := range ctx.Request.Form {
			for _, val := range val {
				arg = append(arg, k, val)
			}
		}
		reflect.Bind(req, arg...)
	} else if err := ctx.Bind(req); err != nil {
		return nil, err
	}
	if valid, ok := req.(interface{ Validate() error }); ok {
		if err := valid.Validate(); err != nil {
			return nil, err
		}
	}
	return req, nil
}
