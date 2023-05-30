package service

import (
	"context"
	"encoding/json"
	"fmt"

	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

func Init(container *dig.Container) {
	container.Provide(NewUserService)
}

func cacheKey(cache repository.Cache, model model.Model) string {
	return fmt.Sprintf("%s:%d", model.Table(), model.GetID())
}
func CacheSet(cache repository.Cache, model model.Model) error {
	if buf, err := json.Marshal(model); err == nil {
		return cache.Set(cacheKey(cache, model), string(buf))
	} else {
		return err
	}
}
func CacheGet(cache repository.Cache, model model.Model) error {
	if buf, err := cache.Get(cacheKey(cache, model)); err == nil {
		return json.Unmarshal([]byte(buf), model)
	} else {
		return err
	}
}
func CacheDel(cache repository.Cache, model model.Model) error {
	return cache.Del(cacheKey(cache, model))
}
func QueueSend(queue repository.Queue, ctx context.Context, op string, model model.Model) error {
	_, err := queue.Send(ctx, model.Table(), op, fmt.Sprintf("%d", model.GetID()))
	return err
}
func QueueRecv(queue repository.Queue, name, topic string, cb func(ctx context.Context, key string, payload string) error) error {
	return queue.Recv(name, topic, cb)
}
