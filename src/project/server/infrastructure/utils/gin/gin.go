package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/router"
)

type Engine struct {
	*gin.Engine
}

func NewEngine(config *config.Config) *Engine {
	if config.Server.Type == enums.Service.HTTP {
		engine := gin.New()
		engine.GET("/metrics", func(ctx *gin.Context) { promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request) })
		return &Engine{engine}
	}
	return nil
}
func (s *Engine) Register(name string, controller interface{}) {
	if s != nil && s.Engine != nil {
		router.Register(s.Engine, name, controller)
	}
}
