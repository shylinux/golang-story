package repository

type Engine interface {
	Insert(obj interface{}) error
	Delete(obj interface{}, id int64) error
	Update(obj interface{}, id int64) error
	SelectOne(obj interface{}, id int64) (interface{}, error)
	SelectList(obj interface{}, res interface{}, page, count int64) error
}

type Cache interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Del(key string) error
}

type Queue interface {
	Send(topic, key string, payload []byte) (string, error)
}
