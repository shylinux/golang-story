package service

import (
	"encoding/json"
	"fmt"
	"reflect"

	"go.uber.org/dig"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

func Init(container *dig.Container) {
	container.Provide(NewUserService)
}
func CacheSet(cache repository.Cache, id int64, obj interface{}) error {
	if buf, err := json.Marshal(obj); err == nil {
		t := reflect.TypeOf(obj)
		return cache.Set(fmt.Sprintf("%s:%d", t.Name(), id), string(buf))
	} else {
		return err
	}
}
func CacheGet(cache repository.Cache, id int64, obj interface{}) error {
	t := reflect.TypeOf(obj)
	if buf, err := cache.Get(fmt.Sprintf("%s:%d", t.Name(), id)); err == nil {
		if err := json.Unmarshal([]byte(buf), obj); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}
func CacheDel(cache repository.Cache, id int64, obj interface{}) error {
	t := reflect.TypeOf(obj)
	return cache.Del(fmt.Sprintf("%s:%d", t.Name(), id))
}
