package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type HttpClient struct {
	Domain string
}

func NewHttpClient(domain string) *HttpClient {
	if domain == "" {
		domain = "http://localhost:8081/"
	}
	return &HttpClient{domain}
}
func (s *HttpClient) Post(uri string, req interface{}, res interface{}) error {
	buf, err := json.Marshal(req)
	if err != nil {
		return err
	}
	log.Printf("post %v %v\n", s.Domain+uri, string(buf))
	if resp, err := http.Post(s.Domain+uri, "application/json", bytes.NewBuffer(buf)); err != nil {
		return err
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf(resp.Status)
		}
		return json.NewDecoder(resp.Body).Decode(res)
	}
}
func (s *HttpClient) Get(uri string, req interface{}, res interface{}) error {
	t := reflect.TypeOf(req).Elem()
	v := reflect.ValueOf(req).Elem()
	for i := 0; i < v.NumField(); i++ {
		var val interface{}
		switch v.Field(i).Kind() {
		case reflect.Int64:
			val = v.Field(i).Int()
		case reflect.String:
			val = v.Field(i).String()
		default:
			continue
		}
		if strings.Contains(uri, "?") {
			uri += "&"
		} else {
			uri += "?"
		}
		uri += fmt.Sprintf("%s=%v", t.Field(i).Name, val)
	}
	log.Printf("get %v\n", s.Domain+uri)
	if resp, err := http.Get(s.Domain + uri); err != nil {
		return err
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf(resp.Status)
		}
		return json.NewDecoder(resp.Body).Decode(res)
	}
}
