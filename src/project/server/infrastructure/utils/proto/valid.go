package proto

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

var validate = validator.New()

func Valid(req interface{}, value interface{}, field string, rule string) error {
	ls := strings.Split(rule, " ")
	switch ls[0] {
	case "default":
		switch v := reflect.ValueOf(req).Elem().FieldByName(Capital(field)); value := value.(type) {
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
			}
			if ls[1] == ">=" && len(value) >= int(length) {
				break
			}
			return errors.NewInvalidParams(fmt.Errorf("%s need %s", field, rule))
		}
	default:
		if err := validate.Var(value, rule); err != nil {
			return errors.NewInvalidParams(fmt.Errorf("%s %s", field, err))
		}
	}
	return nil
}

func (s *Generate) GenValid() {
	s.OpenProto(func(file *os.File, name string) {
		message, comment, action, output := "", "", map[string]template.HTML{}, []string{}
		s.ScanProto(file, func(ls []string, text string) {
			if ls[0] == MESSAGE && strings.HasSuffix(text, "Request {") {
				message, comment, action = ls[1], "", map[string]template.HTML{}
			} else if message == "" {

			} else if strings.HasPrefix(text, "}") {
				if len(action) > 0 {
					output = append(output, s.Template(_valid_template, map[string]interface{}{"message": message, "action": action}))
				}
				message = ""
			} else if strings.HasPrefix(text, "//") {
				comment = text
			} else if comment != "" {
				comment, action[ls[1]] = "", template.HTML(strings.TrimSpace(strings.TrimPrefix(comment, "//")))
			}
		})
		if len(output) == 0 {
			return
		}
		s.Output(path.Join(s.Config.Generate.PbPath, name+"_valid.pb.go"), func(echo func(string, ...interface{})) {
			echo(_valid_import, path.Base(s.Config.Generate.PbPath), path.Dir(logs.ModPath(1)))
			for _, v := range output {
				echo(v)
			}
		})
	})
}

var (
	_valid_import = `
package %s

import %q

`
	_valid_template = `
func (this *{{ .message }}) Validate() error {
{{ range $key, $value := .action }}
	if err := proto.Valid(this, this.{{ $key | Capital }}, "{{ $key }}", "{{ $value }}"); err != nil {
		return err
	}
{{ end }}
	return nil
}
`
)
