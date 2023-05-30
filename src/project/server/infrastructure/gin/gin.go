package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewEngine() *gin.Engine {
	engine := gin.New()
	engine.GET("/metrics", func(ctx *gin.Context) { promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request) })
	return engine
}
