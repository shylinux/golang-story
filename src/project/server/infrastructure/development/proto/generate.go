package proto

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

type Generate struct {
	conf   config.Generate
	protos map[string]map[string]*Item
}

func NewGenerate(config *config.Config, logger logs.Logger, cmds *cmds.Cmds) *Generate {
	s := &Generate{conf: config.Generate}
	cmds.Add("generate", "proto generate", func(ctx context.Context, arg ...string) {
		s.protos = map[string]map[string]*Item{}
		s.OpenProto(func(file *os.File, name string) { s.protos[name] = s.ParseProto(file) })
		s.GenProto()
		s.GenValid()
		s.GenTests()
		s.GenGoAPI()
		s.GenShCLI()
		s.GenJsAPI()
		s.GenErrors()
	})
	return s
}
func (s *Generate) OpenProto(cb func(*os.File, string)) error {
	list, err := system.ReadDir(s.conf.Path)
	if err != nil {
		return errors.New(err, "read proto dir failure")
	}
	for _, info := range list {
		if !strings.HasSuffix(info.Name(), ".proto") {
			continue
		}
		if f, err := system.Open(path.Join(s.conf.Path, info.Name())); err == nil {
			defer f.Close()
			cb(f, strings.TrimSuffix(path.Base(f.Name()), ".proto"))
		}
	}
	return nil
}
func (s *Generate) ScanProto(f *os.File, cb func([]string, string)) {
	for bio := bufio.NewScanner(f); bio.Scan(); {
		text := strings.TrimSpace(bio.Text())
		if text == "" {
			continue
		}
		text = strings.TrimPrefix(text, "repeated")
		text = strings.TrimSuffix(text, ";")
		text = strings.TrimSpace(text)
		cb(strings.Split(text, " "), text)
	}
}

func (s *Generate) Render(name string, tmpl string, data interface{}, funcs template.FuncMap) error {
	f, e := system.Create(name)
	if e != nil {
		return errors.New(e, "render file")
	}
	defer f.Close()
	if e := system.NewTemplate(name, strings.TrimPrefix(tmpl, "\n"), funcs, f, data); e != nil {
		return errors.New(e, "render file")
	} else {
		if strings.HasSuffix(name, ".go") {
			system.Command("", "gofmt", "-w", name)
		}
		return nil
	}
}
func (s *Generate) Template(tmpl string, data interface{}, funcs template.FuncMap) string {
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	system.NewTemplate(errors.FileLine(2), tmpl, funcs, buf, data)
	return string(buf.Bytes())
}
func (s *Generate) Output(name string, cb func(func(str string, arg ...interface{}))) error {
	f, e := system.Create(name)
	if e != nil {
		return errors.New(e, "output file")
	}
	defer f.Close()
	cb(func(str string, arg ...interface{}) {
		fmt.Fprintf(f, strings.TrimSuffix(strings.TrimPrefix(str, "\n"), "\n"), arg...)
		fmt.Fprintln(f)
	})
	return nil
}
func Capital(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToUpper(name[:1]) + name[1:]
}
