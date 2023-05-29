package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type cache struct {
	rdb *redis.Client
}

func (s *cache) Set(key string, value string) error {
	return s.rdb.Set(context.TODO(), key, value, 0).Err()
}
func (s *cache) Get(key string) (string, error) {
	return s.rdb.Get(context.TODO(), key).Result()
}
func (s *cache) Del(key string) error {
	return s.rdb.Del(context.TODO(), key).Err()
}

func New(config *config.Config) (repository.Cache, error) {
	conf := config.Storage.Cache
	return &cache{redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%s", conf.Host, conf.Port), Password: conf.Password})}, nil
}
