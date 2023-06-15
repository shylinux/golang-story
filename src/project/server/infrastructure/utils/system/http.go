package system

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func Request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logs.Errorf("request failure %s %s %s", method, url, err)
		return nil, err
	}
	req.Header.Set("User-Agent", "curl/7.87.0")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logs.Errorf("request failure %s %s %s", method, url, err)
		return nil, err
	}
	return res, nil
}
func RequestJSON(method, url string, data interface{}, res interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if data == nil {
		buf = nil
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	} else {
		switch resp.StatusCode {
		case http.StatusOK, http.StatusCreated:
			if res == nil {
				return nil
			}
			return json.Unmarshal(body, &res)
		default:
			return errors.New(fmt.Errorf(resp.Status), "")
		}
	}
}
