package elasticsearch

import (
	"context"
	"encoding/json"
	"net/http"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/consul"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/repository"
)

type search struct {
	conf config.Search
}

func New(config *config.Config, consul consul.Consul) (repository.Search, error) {
	conf := config.Engine.Search
	if config.Proxy.Simple {
		conf.Enable = false
	}
	if !conf.Enable {
		return &search{conf}, nil
	}
	if list, err := consul.Resolve(config.WithDef(conf.Type, "elasticsearch")); err == nil && len(list) > 0 {
		conf.Host, conf.Port = list[0].Host, list[0].Port
	}
	logs.Infof("engine connect elasticsearch %s:%d/%s", conf.Host, conf.Port, conf.Index)
	return &search{conf}, nil
}
func (s *search) Update(ctx context.Context, mapping string, id int64, data interface{}) error {
	if !s.conf.Enable {
		return nil
	}
	return s.request(ctx, http.MethodPost, s.address("%s/%d", mapping, id), data, nil)
}
func (s *search) Delete(ctx context.Context, mapping string, id int64) error {
	if !s.conf.Enable {
		return nil
	}
	return s.request(ctx, http.MethodDelete, s.address("%s/%d", mapping, id), nil, nil)
}
func (s *search) Query(ctx context.Context, mapping string, res interface{}, page, count int64, key, value string) (total int64, err error) {
	if !s.conf.Enable {
		return 0, nil
	}
	var data struct {
		Hits struct {
			Total struct{ Value int64 }
			Hits  []struct {
				Index  string                 `json:"_index"`
				Type   string                 `json:"_type"`
				ID     string                 `json:"_id"`
				Score  float64                `json:"_score"`
				Source map[string]interface{} `json:"_source"`
			}
		}
	}
	if err := s.request(ctx, http.MethodGet, s.address("%s/_search?q=%s:%s", mapping, key, value), nil, &data); err != nil {
		return 0, err
	}
	list := []map[string]interface{}{}
	for _, v := range data.Hits.Hits {
		list = append(list, v.Source)
	}
	buf, err := json.Marshal(list)
	if err == nil {
		err = json.Unmarshal(buf, res)
	}
	return data.Hits.Total.Value, err
}
