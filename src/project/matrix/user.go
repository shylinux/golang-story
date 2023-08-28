package matrix

import (
	"shylinux.com/x/ice"
)

type user struct {
	Client
	idname string `data:"userID"`
	create string `name:"create username* password* email"`
	list   string `name:"list userID auto" help:"用户"`
}

func init() { ice.CodeCtxCmd(user{}) }
