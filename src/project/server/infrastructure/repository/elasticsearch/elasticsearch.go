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
	*config.Config
}

func New(config *config.Config, consul consul.Consul) (repository.Search, error) {
	conf := config.Engine.Search
	if list, err := consul.Resolve(config.WithDef(conf.Name, "elasticsearch")); err == nil && len(list) > 0 {
		conf.Host, conf.Port = list[0].Host, list[0].Port
	}
	logs.Infof("engine connect elasticsearch %s:%d/%s", conf.Host, conf.Port, conf.Index)
	return &search{config}, nil
}
func (s *search) Update(ctx context.Context, mapping string, id int64, data interface{}) error {
	return s.request(ctx, http.MethodPost, s.address("%s/%d", mapping, id), data, nil)
}
func (s *search) Delete(ctx context.Context, mapping string, id int64) error {
	return s.request(ctx, http.MethodDelete, s.address("%s/%d", mapping, id), nil, nil)
}
func (s *search) Query(ctx context.Context, mapping string, res interface{}, page, count int64, key, value string) (total int64, err error) {
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
