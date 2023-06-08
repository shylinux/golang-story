package repository

import "context"

type Storage interface {
	Insert(ctx context.Context, obj interface{}) error
	Delete(ctx context.Context, obj interface{}, id int64) error
	Update(ctx context.Context, obj interface{}, id int64) error
	SelectOne(ctx context.Context, obj interface{}, id int64) (interface{}, error)
	SelectList(ctx context.Context, obj interface{}, res interface{}, page, count int64, condition string, arg ...interface{}) (int64, error)
	AutoMigrate(obj ...interface{}) error
}

type Cache interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Del(key string) error
}

type Queue interface {
	Send(ctx context.Context, topic, key string, payload []byte) (string, error)
	Recv(ctx context.Context, name, topic string, cb func(ctx context.Context, key string, payload []byte) error) error
}
