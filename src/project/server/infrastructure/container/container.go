package container

import (
	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/check"
)

type Container struct {
	*dig.Container
}

func (s *Container) Invoke(cb interface{}) {
	check.Assert(s.Container.Invoke(cb))
}

func New(cb ...func(*Container)) *Container {
	container := &Container{dig.New()}
	for _, cb := range cb {
		cb(container)
	}
	return container
}
