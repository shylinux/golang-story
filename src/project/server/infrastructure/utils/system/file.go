package system

import (
	"html/template"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func Create(name string) (*os.File, error) {
	if _, e := os.Stat(path.Dir(name)); os.IsNotExist(e) {
		os.MkdirAll(path.Dir(name), 0755)
	}
	if f, e := os.Create(name); e != nil {
		logs.Errorf("file create %s failure %s", name, e)
		return nil, errors.New(e, "file create failure")
	} else {
		logs.Infof("file create %s", name)
		return f, nil
	}
}
func Open(name string) (*os.File, error) {
	if f, err := os.Open(name); err != nil {
		logs.Errorf("file open %s failure %s", name, err)
		return nil, err
	} else {
		logs.Infof("file open %s", name)
		return f, nil
	}
}
func ReadDir(name string) ([]os.DirEntry, error) {
	if list, err := os.ReadDir(name); err != nil {
		logs.Errorf("dir read %s failure %s", name, err)
		return nil, err
	} else {
		logs.Infof("dir read %s", name)
		return list, nil
	}
}
func NewTemplate(name string, tmpl string, funcs template.FuncMap, f io.Writer, data interface{}) error {
	if t, e := template.New(name).Funcs(funcs).Parse(tmpl); e != nil {
		logs.Errorf("file render %s %s %s:%s", name, e, errors.FileLine(2), errors.FuncName(2))
		return e
	} else if e := t.Execute(f, data); e != nil {
		logs.Errorf("file render %s %s %s %s %s:%s", name, e, f, data, errors.FileLine(2), errors.FuncName(2))
		return e
	} else {
		return nil
	}
}
func WriteFile(name string, data []byte, perm fs.FileMode) error {
	if _, e := os.Stat(path.Dir(name)); os.IsNotExist(e) {
		os.MkdirAll(path.Dir(name), 0755)
	}
	if err := ioutil.WriteFile(name, data, perm); err != nil {
		logs.Errorf("file write %s %s", name, err)
		return err
	}
	return nil
}
