package deploy

import (
	"archive/zip"
	"context"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/cmds"
)

type Deploy struct {
	*config.Config
}

func New(cmds *cmds.Cmds, config *config.Config, logger logs.Logger) *Deploy {
	deploy := &Deploy{config}
	cmds.Add("deploy", "deploy", func(ctx context.Context, arg ...string) {
		// deploy.Install("consul")
		deploy.Unpack("consul")
	})
	return deploy
}

func (s *Deploy) Download() {
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
	res, err := http.Get(target.Address)
	if err != nil {
		logs.Errorf("download failure %s %s %s", name, target.Address, err)
		return
	}
	defer res.Body.Close()
	f, e := os.Create(_target)
	if e != nil {
		logs.Errorf("download failure %s %s %s", name, target.Address, e)
		return
	}
	defer f.Close()
	length, _ := strconv.Atoi(res.Header.Get("Content-Length"))
	total, last := 0, 0
	logs.Infof("download start %s %s %s", name, target.Address, logs.Size(int64(length)))
	if n, e := io.Copy(io.MultiWriter(f, newprocess(func(buf []byte) {
		if total += len(buf); total/length != last {
			logs.Infof("download process %s %s %s %s %s", name, target.Address, logs.Size(int64(total)), logs.Percent(int64(total), int64(length)), logs.Cost(begin))
			last = total / length
		}
	})), res.Body); e != nil {
		logs.Errorf("download failure %s %s %s", name, target.Address, e)
	} else {
		logs.Infof("download success %s %s %s %s", name, target.Address, logs.Size(n), logs.Cost(begin))
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
func (s *Deploy) Unpack(name string) {
	target := s.Config.Install.GetTarget(name)
	if strings.HasSuffix(target.Address, ".zip") {
		s.UnpackZIP(name)
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
