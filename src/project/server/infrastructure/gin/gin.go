package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"shylinux.com/x/golang-story/src/project/server/domain/enums"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
)

func NewEngine(config *config.Config) *gin.Engine {
	if config.Service.Type == enums.Service.HTTP {
		engine := gin.New()
		engine.GET("/metrics", func(ctx *gin.Context) { promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request) })
		return engine
	}
	return nil
}
