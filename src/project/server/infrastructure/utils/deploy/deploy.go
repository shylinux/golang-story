package deploy

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/cmds"
)

type Deploy struct {
	*config.Config
}

func New(cmds *cmds.Cmds, config *config.Config, logger logs.Logger) *Deploy {
	s := &Deploy{config}
	cmds.Add("deploy", "deploy", func(ctx context.Context, arg ...string) {
		// deploy.Install("consul")
		s.Unpack("consul")
	})
	cmds.Add("install", "install", func(ctx context.Context, arg ...string) {
		if len(arg) == 0 {
			buf, _ := json.MarshalIndent(config.Install, "", "  ")
			fmt.Println(string(buf))
			return
		}
		s.Install(arg[0])
	})
	cmds.Add("unpack", "unpack", func(ctx context.Context, arg ...string) {
		if len(arg) == 0 {
			buf, _ := json.MarshalIndent(config.Install, "", "  ")
			fmt.Println(string(buf))
			return
		}
		s.Unpack(arg[0])
	})
	return s
}

func (s *Deploy) Create(name string, mode os.FileMode) (*os.File, error) {
	os.MkdirAll(path.Dir(name), 0755)
	os.OpenFile(name, os.O_CREATE|os.O_WRONLY, mode)
	return os.Create(name)
}
func (s *Deploy) Install(name string) {
	target := s.Config.Install.GetTarget(name)
	logs.Infof("download prepare %s %s", name, target.Address)
	_target := path.Join("usr", path.Base(target.Address))
	if s, e := os.Stat(_target); e == nil {
		logs.Infof("download success %s %s %s", name, target.Address, logs.Size(s.Size()))
		return
	}
	begin := time.Now()
	req, err := http.NewRequest(http.MethodGet, target.Address, nil)
	if err != nil {
		logs.Errorf("download failure %s %s %s", name, target.Address, err)
		return
	}
	req.Header.Set("User-Agent", "curl/7.87.0")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logs.Errorf("download failure %s %s %s", name, target.Address, err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		logs.Errorf("download failure %s", res.Status)
		return
	}
	os.MkdirAll(path.Dir(_target), 0755)
	f, e := os.Create(_target)
	if e != nil {
		logs.Errorf("download failure %s %s %s", name, target.Address, e)
		return
	}
	defer f.Close()
	length, _ := strconv.Atoi(res.Header.Get("Content-Length"))
	if length == 0 {
		length = -1
	}
	logs.Infof("download start %s %s %s", name, target.Address, logs.Size(int64(length)))
	if n, e := io.Copy(io.MultiWriter(f, progressbar.DefaultBytes(res.ContentLength, "正在下载")), res.Body); e != nil {
		logs.Errorf("download failure %s %s %s", name, target.Address, e)
	} else {
		logs.Infof("download success %s %s %s %s", name, target.Address, logs.Size(n), logs.Cost(begin))
	}
}
func (s *Deploy) Unpack(name string) {
	target := s.Config.Install.GetTarget(name)
	if strings.HasSuffix(target.Address, ".zip") {
		s.UnpackZIP(name)
	} else if strings.HasSuffix(target.Address, ".tar.gz") {
		s.UnpackGZIP(name)
	}
}
func (s *Deploy) UnpackGZIP(name string) {
	target := s.Config.Install.GetTarget(name)
	r, e := os.Open(path.Join("usr", path.Base(target.Address)))
	if e != nil {
		logs.Errorf("unpack failure %s %s", name, e)
		return
	}
	g, e := gzip.NewReader(r)
	if e != nil {
		logs.Errorf("unpack failure %s %s", name, e)
		return
	}
	t := tar.NewReader(g)
	count := 0
	for {
		_, e := t.Next()
		if e != nil {
			break
		}
		count++
	}
	r.Seek(0, 0)
	g, e = gzip.NewReader(r)
	if e != nil {
		logs.Errorf("unpack failure %s %s", name, e)
		return
	}
	t = tar.NewReader(g)
	bar := progressbar.Default(int64(count), "正在解压")
	for i := 0; i < count; i++ {
		h, e := t.Next()
		if e != nil {
			logs.Errorf("unpack failure %s %s", h.Name, e)
			break
		}
		bar.Add(1)
		_name := path.Base(h.Name)
		if len(_name) < 10 {
			_name += strings.Repeat(" ", 10-len(_name))
		} else {
			_name = _name[:10]
		}

		bar.Describe(_name)
		if h.FileInfo().IsDir() {
			os.MkdirAll(path.Dir(name), h.FileInfo().Mode())
			continue
		}
		func() {
			f, e := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, h.FileInfo().Mode())
			if e != nil {
				logs.Errorf("unpack failure %s %s", name, e)
				return
			}
			io.Copy(f, t)
			defer f.Close()
		}()
	}
}
func (s *Deploy) UnpackZIP(name string) {
	target := s.Config.Install.GetTarget(name)
	r, e := zip.OpenReader(path.Join("usr", path.Base(target.Address)))
	if e != nil {
		logs.Errorf("unpack failure %s %s", name, e)
		return
	}
	for _, file := range r.File {
		func() {
			r, e := file.Open()
			if e != nil {
				logs.Errorf("unpack failure %s %s", file.Name, e)
				return
			}
			defer r.Close()
			w, e := os.OpenFile(file.Name, os.O_CREATE|os.O_WRONLY, file.Mode())
			if e != nil {
				logs.Errorf("unpack failure %s %s", file.Name, e)
				return
			}
			defer w.Close()
			n, e := io.Copy(w, r)
			if e != nil {
				logs.Errorf("unpack failure %s %s", file.Name, e)
				return
			}
			logs.Infof("unpack %s %s", file.Name, logs.Size(n))
		}()
	}
}
func (s *Deploy) Clone() {
}
func (s *Deploy) Build() {
}
func (s *Deploy) Start() {
}

type process struct {
	cb func([]byte)
}

func newprocess(cb func([]byte)) *process {
	return &process{cb}
}
func (s process) Write(buf []byte) (n int, e error) {
	s.cb(buf)
	return len(buf), nil
}
