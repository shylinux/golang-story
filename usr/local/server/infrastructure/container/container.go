package container

import (
	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
)

type Container struct {
	*dig.Container
}

func New(cb ...func(*Container)) *Container {
	container := &Container{dig.New()}
	container.Provide(func() *Container { return container })
	return container.Add(cb...)
}
func (s *Container) Add(cb ...func(*Container)) *Container {
	for _, cb := range cb {
		cb(s)
	}
	return s
}
func (s *Container) Provide(cb ...interface{}) *Container {
	for _, cb := range cb {
		s.Container.Provide(cb)
	}
	return s
}
func (s *Container) Invoke(cb interface{}) {
	errors.Assert(s.Container.Invoke(cb))
}
