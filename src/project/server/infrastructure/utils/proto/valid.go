package proto

import (
	"fmt"
	"html/template"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
)

var validate = validator.New()

func Valid(req interface{}, value interface{}, name string, rule string) error {
	switch ls := strings.Split(rule, " "); ls[0] {
	case "default":
		switch v := reflect.ValueOf(req).Elem().FieldByName(Capital(name)); value := value.(type) {
		case int64:
			if value > 0 {
				break
			} else if val, err := strconv.ParseInt(ls[1], 10, 64); err == nil {
				v.SetInt(val)
			}
		case string:
			if value != "" {
				break
			}
			v.SetString(ls[1])
		}
	case "length":
		switch value := value.(type) {
		case string:
			length, _ := strconv.ParseInt(ls[2], 10, 64)
			if ls[1] == ">" && len(value) > int(length) {
				break
			} else if ls[1] == ">=" && len(value) >= int(length) {
				break
			}
			return errors.NewInvalidParams(fmt.Errorf("%s need %s", name, rule))
		}
	default:
		if err := validate.Var(value, rule); err != nil {
			return errors.NewInvalidParams(fmt.Errorf("%s %s", name, err))
		}
	}
	return nil
}

func (s *Generate) GenValid() {
	for name, proto := range s.protos {
		s.Render(path.Join(s.conf.PbPath, name+"_valid.pb.go"), _valid_template, proto, template.FuncMap{
			"RequestList": func() (res []string) {
				for _, service := range proto[PACKAGE].List {
					for _, method := range proto[service].List {
						res = append(res, proto[method].List[0])
					}
				}
				return
			},
			"RequestField": func(request string) (res []map[string]interface{}) {
				list := proto[request].List
				for i := 0; i < len(list); i += 3 {
					if list[i] == "" {
						continue
					}
					res = append(res, map[string]interface{}{
						"field": Capital(list[i+2]),
						"rule":  template.HTML(list[i]),
						"name":  list[i+2],
					})
				}
				return
			},
		})
	}
}

var (
	_valid_template = `
package pb

import "shylinux.com/x/golang-story/src/project/server/infrastructure/utils/proto"

{{ range $index, $request := RequestList }}
func (this *{{ $request }}) Validate() error {
{{ range $index, $field := RequestField $request }}
	if err := proto.Valid(this, this.{{ index $field "field" }}, "{{ index $field "name" }}", "{{ index $field "rule" }}"); err != nil {
		return err
	}
{{ end }}
	return nil
}
{{ end }}
`
)
