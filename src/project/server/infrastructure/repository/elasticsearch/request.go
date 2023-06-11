package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func (s *search) address(str string, arg ...interface{}) string {
	conf := s.Config.Engine.Search
	return fmt.Sprintf("http://%s:%d/%s/", conf.Host, conf.Port, conf.Index) + fmt.Sprintf(str, arg...)
}
func (s *search) request(ctx context.Context, method string, url string, data interface{}) (body []byte, err error) {
	begin := time.Now()
	logs.Infof("search %s request %s %s", method, url, logs.Marshal(data), ctx)
	echo := func(res []byte, err error) ([]byte, error) {
		if err != nil && err.Error() != "" {
			logs.Warnf("search %s response %s %s %s cost:%s", method, url, err, string(res), logs.Cost(begin), ctx)
		} else {
			logs.Infof("search %s response %s %s cost:%s", method, url, string(res), logs.Cost(begin), ctx)
		}
		return res, err
	}
	buf, err := json.Marshal(data)
	if err != nil {
		return echo(nil, errors.New(err, "search "+method))
	}
	if data == nil {
		buf = nil
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(buf))
	if err != nil {
		return echo(nil, errors.New(err, "search "+method))
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return echo(nil, errors.New(err, "search "+method))
	}
	defer res.Body.Close()
	if body, err := ioutil.ReadAll(res.Body); err != nil {
		return echo(body, errors.New(err, "search "+method))
	} else {
		switch res.StatusCode {
		case http.StatusOK, http.StatusCreated:
			return echo(body, nil)
		default:
			return echo(body, errors.New(fmt.Errorf(res.Status), "search "+method))
		}
	}
}
