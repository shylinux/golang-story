package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type cache struct {
	conf config.Cache
	*redis.Client
}

func New(config *config.Config, consul consul.Consul) repository.Cache {
	conf := config.Engine.Cache
	if config.Proxy.Simple {
		conf.Enable = false
	}
	if !conf.Enable {
		return &cache{conf, nil}
	}
	if list, err := consul.Resolve(config.WithDef(conf.Name, "redis")); err == nil && len(list) > 0 {
		conf.Host, conf.Port = list[0].Host, list[0].Port
	}
	logs.Infof("engine connect redis %s:%d", conf.Host, conf.Port)
	return &cache{conf, redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%d", conf.Host, conf.Port), Password: conf.Password})}
}
func (s *cache) Set(ctx context.Context, key string, value string) error {
	if !s.conf.Enable {
		return nil
	}
	if err := s.Client.Set(ctx, key, value, 0).Err(); err != nil {
		logs.Warnf("redis set %s %s %s", key, value, err, ctx)
		return errors.New(err, "redis set")
	} else {
		logs.Infof("redis set %s %s", key, value, ctx)
		return nil
	}
}
func (s *cache) Get(ctx context.Context, key string) (string, error) {
	if !s.conf.Enable {
		return "", errors.New(fmt.Errorf("not found"), "")
	}
	if res, err := s.Client.Get(ctx, key).Result(); err != nil {
		logs.Warnf("redis get %s %s", key, err, ctx)
		return res, errors.New(err, "redis get")
	} else {
		logs.Infof("redis get %s %s", key, res, ctx)
		return res, nil
	}
}
func (s *cache) Del(ctx context.Context, key string) error {
	if !s.conf.Enable {
		return nil
	}
	if err := s.Client.Del(ctx, key).Err(); err != nil {
		logs.Warnf("redis del %s %s", key, err, ctx)
		return errors.New(err, "redis del")
	} else {
		logs.Infof("redis del %s", key, ctx)
		return nil
	}
}
