package service

import (
	"github.com/gin-gonic/gin"
	"shylinux.com/x/golang-story/src/project/server/application"
)

type AdminService struct {
	app *application.AdminApp
}

func NewAdminService(app *application.AdminApp, engine *gin.Engine) *AdminService {
	register(engine, "admin", app)
	return &AdminService{app}
}
