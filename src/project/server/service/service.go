package service

import (
	"context"
	"encoding/json"
	"fmt"

	"shylinux.com/x/golang-story/src/project/server/domain/model"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/container"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

func Init(container *container.Container) {
	container.Provide(NewAuthService)
	container.Provide(NewUserService)
}

func Clause(cond bool, stmt string, arg ...interface{}) (string, []interface{}) {
	if cond {
		return stmt, arg
	}
	return "", []interface{}{}
}
func SearchUpdate(ctx context.Context, search repository.Search, mapping string, id int64, data interface{}) error {
	if err := search.Update(ctx, mapping, id, data); err != nil {
		logs.Errorf(errors.New(err, "search update failure %s", logs.FileLine(2)).Error(), ctx)
		return err
	}
	return nil
}
func SearchDelete(ctx context.Context, search repository.Search, mapping string, id int64) error {
	if err := search.Delete(ctx, mapping, id); err != nil {
		logs.Warnf(errors.New(err, "search delete failure %s", logs.FileLine(2)).Error(), ctx)
		return err
	}
	return nil
}
func QueueSend(ctx context.Context, queue repository.Queue, topic, op string, v interface{}) error {
	if buf, err := json.Marshal(v); err != nil {
		logs.Errorf(errors.New(err, "message send failure topic: %s operate: %s payload: %+s %s", topic, op, v, logs.FileLine(2)).Error(), ctx)
		return err
	} else if _, err = queue.Send(ctx, topic, op, buf); err != nil {
		logs.Errorf(errors.New(err, "message send failure topic: %s operate: %s payload: %+s %s", topic, op, v, logs.FileLine(2)).Error(), ctx)
		return err
	}
	return nil
}
func CacheSet(ctx context.Context, cache repository.Cache, prefix string, model model.Model) error {
	if buf, err := json.Marshal(model); err != nil {
		logs.Warnf(errors.New(err, "cache set failure %s", logs.FileLine(2)).Error(), ctx)
		return err
	} else if err := cache.Set(ctx, cacheKey(cache, prefix, model), string(buf)); err != nil {
		logs.Warnf(errors.New(err, "cache set failure %s", logs.FileLine(2)).Error(), ctx)
		return err
	}
	return nil
}
func CacheDel(ctx context.Context, cache repository.Cache, prefix string, model model.Model) error {
	if err := cache.Del(ctx, cacheKey(cache, prefix, model)); err != nil {
		logs.Warnf(errors.New(err, "cache del failure %s", logs.FileLine(2)).Error(), ctx)
		return err
	}
	return nil
}
func CacheGet(ctx context.Context, cache repository.Cache, prefix string, model model.Model) error {
	if buf, err := cache.Get(ctx, cacheKey(cache, prefix, model)); err != nil {
		logs.Warnf(errors.New(err, "cache get failure %s", logs.FileLine(2)).Error(), ctx)
		return err
	} else if err := json.Unmarshal([]byte(buf), model); err != nil {
		logs.Warnf(errors.New(err, "cache get failure %s", logs.FileLine(2)).Error(), ctx)
		return err
	}
	return nil
}
func cacheKey(cache repository.Cache, prefix string, model model.Model) string {
	return fmt.Sprintf("%s:%s", prefix, model.GetID())
}
