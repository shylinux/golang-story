package application

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/domain"
)

type AdminApp struct {
}

func NewAdminApp() *AdminApp {
	return &AdminApp{}
}

func (app *AdminApp) List(ctx context.Context) ([]*domain.Admin, error) {
	return []*domain.Admin{&domain.Admin{Name: "shy"}}, nil
}
