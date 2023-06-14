package context

import "context"

type Context struct {
	context.Context
}

func New() *Context {
	return &Context{context.Background()}
}

func (s *Context) Context() {
	return s.Context
}
