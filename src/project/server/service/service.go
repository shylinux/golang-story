package service

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gogather/com/log"
	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/response"
)

type MainService struct {
	admin *AdminService
	*gin.Engine
}

func NewMainService(admin *AdminService, engine *gin.Engine) *MainService {
	return &MainService{admin, engine}
}

func Init(container *dig.Container) {
	container.Provide(gin.New)
}

// register 注册服务
func register(g *gin.Engine, group string, s interface{}) {
	grp := g.Group(group)

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumMethod(); i++ {
		grp.POST(strings.ToLower(t.Method(i).Name), wrap(v.Method(i).Interface()))
	}
}

// wrap 包装回调函数
func wrap(handle interface{}) func(*gin.Context) {
	return func(ctx *gin.Context) {
		t := reflect.TypeOf(handle)
		v := reflect.ValueOf(handle)

		var res []reflect.Value
		switch t.NumIn() {
		case 1:
			log.Infof("access %v", ctx.Request.URL)
			res = v.Call([]reflect.Value{reflect.ValueOf(ctx)})
		case 2:
			arg := reflect.New(t.In(1).Elem()).Interface()
			if err := ctx.Bind(arg); err != nil {
				response.WriteParamInvalid(ctx, err)
				return
			}
			log.Infof("access %v %#v", ctx.Request.URL, arg)
			res = v.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(arg)})
		default:
			response.WriteBase(ctx, fmt.Errorf("func arg must be: (ctx, [data])"))
		}

		if len(res) == 0 {
			return
		}

		switch err := res[len(res)-1].Interface().(type) {
		case nil:
			if len(res) == 1 {
				response.WriteBase(ctx, err)
				return
			}
		case error:
			if err != nil {
				response.WriteBase(ctx, err)
				return
			}
		default:
			response.WriteBase(ctx, fmt.Errorf("last res must be error %v", err))
			return
		}

		if len(res) == 2 {
			switch id := res[0].Interface().(type) {
			case int64:
				response.WriteBaseID(ctx, id)
			default:
				response.WriteData(ctx, res[0].Interface(), res[1].Interface())
			}
			return
		}

		if len(res) == 3 {
			switch secondRes := res[1].Interface().(type) {
			case int64:
				response.WriteBasePage(ctx, res[0].Interface(), secondRes)
			default:
				response.WriteBase(ctx, fmt.Errorf("res must be (list [], total int64, err error)"))
			}
			return
		}

		response.WriteBase(ctx, fmt.Errorf("unknown res"))
	}
}
