package system

import (
	"html/template"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func MkdirAll(dir ...string) {
	os.MkdirAll(path.Join(dir...), 0755)
}
func AbsPath(dir string) string {
	if strings.HasPrefix(dir, "/") {
		return dir
	}
	pwd, _ := os.Getwd()
	return path.Join(pwd, dir)
}
func PwdPath(dir string) string {
	if !strings.HasPrefix(dir, "/") {
		return dir
	}
	pwd, _ := os.Getwd()
	return strings.TrimPrefix(dir, pwd+"/")
}
func Exists(name string) bool {
	if _, e := os.Stat(name); e == nil {
		return true
	}
	return false
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
func Remove(name string) error {
	logs.Infof("file remove %s", name)
	os.Remove(name)
	return nil
}
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
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	if _, e := os.Stat(path.Dir(name)); os.IsNotExist(e) {
		os.MkdirAll(path.Dir(name), 0755)
	}
	if f, err := os.OpenFile(name, flag, perm); err != nil {
		logs.Errorf("file open %s failure %s", name, err)
		return nil, err
	} else {
		return f, nil
	}
}
func ReadFile(name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(name)
	return buf, err
}
func WriteFile(name string, data []byte, perm fs.FileMode) error {
	if _, e := os.Stat(path.Dir(name)); os.IsNotExist(e) {
		os.MkdirAll(path.Dir(name), 0755)
	}
	if err := ioutil.WriteFile(name, data, perm); err != nil {
		logs.Errorf("file write %s %s", name, err)
		return err
	} else {
		logs.Infof("file write %s %s", name, string(data))
		return nil
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
func NewTemplateFile(name string, tmpl string, funcs template.FuncMap, data interface{}) error {
	f, e := Create(name)
	if e != nil {
		return errors.New(e, "render file")
	}
	defer f.Close()
	if e := NewTemplate(name, strings.TrimPrefix(tmpl, "\n"), funcs, f, data); e != nil {
		return errors.New(e, "render file")
	}
	return nil
}
