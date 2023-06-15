package deploy

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/schollz/progressbar/v3"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

func (s *Deploy) Unpack(name string) error {
	if system.Exists(s.BinPath(name)) {
		return nil
	} else if p := s.Path(name); strings.HasSuffix(p, ".tar.xz") {
		_, err := system.Command("", "tar", "xf", p, "-C", path.Dir(p))
		return err
	} else if strings.HasSuffix(p, ".tar.gz") {
		return s.UnpackGZIP(name)
	} else if strings.HasSuffix(p, ".zip") {
		return s.UnpackZIP(name)
	} else {
		return errors.New(fmt.Errorf("not implement unpack"), p)
	}
}
func (s *Deploy) UnpackGZIP(name string) error {
	count, err := s.UnpackGZIPCount(name)
	if err != nil {
		return err
	}
	f, e := system.Open(s.Path(name))
	if e != nil {
		logs.Errorf("unpack failure %s %s", name, e)
		return e
	}
	defer f.Close()
	g, e := gzip.NewReader(f)
	if e != nil {
		logs.Errorf("unpack failure %s %s", name, e)
		return e
	}
	t := tar.NewReader(g)
	target := s.Config.Install.GetTarget(name)
	bar := progressbar.Default(int64(count), "正在解压: "+path.Base(target.Address))
	for i := 0; i < count; i++ {
		if h, e := t.Next(); e != nil {
			logs.Errorf("unpack failure %s %s", h.Name, e)
		} else if bar.Add(1); h.FileInfo().IsDir() {
			os.MkdirAll(path.Join(USR, target.Unpack, h.Name), h.FileInfo().Mode())
		} else {
			s.UnpackFile(target, h.Name, h.FileInfo().Mode(), t)
		}
	}
	return nil
}
func (s *Deploy) UnpackGZIPCount(name string) (int, error) {
	f, e := system.Open(s.Path(name))
	if e != nil {
		logs.Errorf("unpack failure %s %s", name, e)
		return 0, e
	}
	defer f.Close()
	g, e := gzip.NewReader(f)
	if e != nil {
		logs.Errorf("unpack failure %s %s", name, e)
		return 0, e
	}
	t := tar.NewReader(g)
	count := 0
	for {
		if _, e := t.Next(); e != nil {
			break
		}
		count++
	}
	return count, nil
}
func (s *Deploy) UnpackZIP(name string) error {
	r, e := zip.OpenReader(s.Path(name))
	if e != nil {
		logs.Errorf("unpack failure %s %s", name, e)
		return e
	}
	target := s.Config.Install.GetTarget(name)
	bar := progressbar.Default(int64(len(r.File)), "正在解压: "+path.Base(target.Address))
	for _, file := range r.File {
		if bar.Add(1); file.FileInfo().IsDir() {
			os.MkdirAll(path.Join(USR, target.Unpack, file.Name), file.FileInfo().Mode())
		} else if f, e := file.Open(); e != nil {
			logs.Errorf("unpack failure %s %s", file.Name, e)
		} else {
			s.UnpackFile(target, file.Name, file.Mode(), f)
		}
	}
	return nil
}
func (s *Deploy) UnpackFile(target config.Target, name string, perm os.FileMode, f io.Reader) error {
	if f, ok := f.(io.Closer); ok {
		defer f.Close()
	}
	w, e := system.OpenFile(path.Join(USR, target.Unpack, name), os.O_CREATE|os.O_WRONLY, perm)
	if e != nil {
		return e
	}
	defer w.Close()
	if _, e := io.Copy(w, f); e != nil {
		logs.Errorf("unpack failure %s %s", name, e)
		return e
	} else {
		return nil
	}
}
