package proto

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

type Generate struct {
	*config.Config
}

func NewGenerate(config *config.Config, logger logs.Logger) *Generate {
	return &Generate{config}
}
func (s *Generate) OpenProto(cb func(*os.File, string)) error {
	list, err := os.ReadDir(s.Config.Generate.Path)
	if err != nil {
		logs.Errorf("read proto dir failure %s", err)
		return errors.New(err, "read proto dir failure")
	}
	for _, info := range list {
		if !strings.HasSuffix(info.Name(), ".proto") {

		} else if file, err := os.Open(path.Join(s.Config.Generate.Path, info.Name())); err != nil {
			logs.Errorf("open proto file failure %s", err)
		} else {
			defer file.Close()
			logs.Infof("open proto file %s", file.Name())
			cb(file, strings.TrimSuffix(path.Base(file.Name()), ".proto"))
		}
	}
	return nil
}
func (s *Generate) ScanProto(file *os.File, cb func([]string, string)) {
	for bio := bufio.NewScanner(file); bio.Scan(); {
		text := strings.TrimSpace(bio.Text())
		if text == "" {
			continue
		}
		cb(strings.Split(text, " "), text)
	}
}
func (s *Generate) Render(file string, tmpl string, data interface{}) error {
	if _, e := os.Stat(path.Dir(file)); os.IsNotExist(e) {
		os.MkdirAll(path.Dir(file), 0755)
	}
	f, e := os.Create(file)
	if e != nil {
		logs.Errorf("  render file %s", e)
		return errors.New(e, "render file")
	}
	defer f.Close()
	if t, e := template.New("render").Parse(strings.TrimPrefix(tmpl, "\n")); e != nil {
		logs.Errorf("  render file %s", e)
		return errors.New(e, "render file")
	} else if e := t.Execute(f, data); e != nil {
		logs.Errorf("  render file %s", e)
		return errors.New(e, "render file")
	} else {
		logs.Infof("  render file %s", file)
		return nil
	}
}
func (s *Generate) Template(tmpl string, data interface{}) string {
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	template.Must(template.New("render").Funcs(template.FuncMap{"Capital": Capital}).Parse(tmpl)).Execute(buf, data)
	return string(buf.Bytes())
}
func (s *Generate) Output(file string, cb func(func(str string, arg ...interface{}))) {
	f, err := os.Create(file)
	if err != nil {
		logs.Errorf("generate %s %s", file, err)
		return
	}
	defer f.Close()
	logs.Infof("  generate %v", file)
	cb(func(str string, arg ...interface{}) {
		fmt.Fprintf(f, strings.TrimSuffix(strings.TrimPrefix(str, "\n"), "\n"), arg...)
		fmt.Fprintln(f)
	})
}
func Capital(field string) string {
	return strings.ToUpper(field[:1]) + field[1:]
}

const (
	PACKAGE = "package"
	SERVICE = "service"
	MESSAGE = "message"
)
