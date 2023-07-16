package repository

import (
	"context"

	"shylinux.com/x/golang-story/src/project/server/domain/model"
)

type Storage interface {
	Insert(ctx context.Context, obj model.Model) error
	Delete(ctx context.Context, obj model.Model) error
	Update(ctx context.Context, obj model.Model) error
	SelectOne(ctx context.Context, obj model.Model) error
	SelectList(ctx context.Context, obj model.Model, res interface{}, page, count int64, condition string, arg ...interface{}) (int64, error)
	AutoMigrate(obj ...interface{}) error
}
type Search interface {
	Update(ctx context.Context, mapping string, id int64, data interface{}) error
	Delete(ctx context.Context, mapping string, id int64) error
	Query(ctx context.Context, mapping string, res interface{}, page, count int64, key, value string) (int64, error)
}
type Cache interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}
type Queue interface {
	Send(ctx context.Context, topic, key string, payload []byte) (string, error)
	Recv(ctx context.Context, name, topic string, cb func(ctx context.Context, key string, payload []byte)) error
}
