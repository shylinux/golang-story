package container

import (
	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
)

type Container struct {
	*dig.Container
}

func (s *Container) Add(cb ...func(*Container)) *Container {
	for _, cb := range cb {
		cb(s)
	}
	return s
}
func (s *Container) Invoke(cb interface{}) {
	errors.Assert(s.Container.Invoke(cb))
}

func New(cb ...func(*Container)) *Container {
	container := &Container{dig.New()}
	container.Provide(func() *Container { return container })
	return container.Add(cb...)
}
