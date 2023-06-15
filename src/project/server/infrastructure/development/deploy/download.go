package deploy

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

func (s *Deploy) Download(name string) error {
	if system.Exists(s.Path(name)) {
		return nil
	}
	target := s.Config.Install.GetTarget(name)
	logs.Infof("download prepare %s %s", name, target.Address)
	if address, ok := s.Config.ReplaceMap[target.Address]; ok {
		logs.Infof("download replace %s => %s", target.Address, address)
		target.Address = address
	} else if address, ok := s.Config.ReplaceMap[path.Base(target.Address)]; ok {
		logs.Infof("download replace %s => %s", target.Address, address)
		target.Address = address
	}
	res, err := system.Request(http.MethodGet, target.Address, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		logs.Errorf("download failure %s", res.Status)
		return errors.New(fmt.Errorf(res.Status), "")
	}
	f, e := system.Create(s.Path(name))
	if e != nil {
		return e
	}
	defer f.Close()
	begin := time.Now()
	logs.Infof("download start %s %s %s", name, target.Address, logs.Size(res.ContentLength))
	if n, e := io.Copy(io.MultiWriter(f, progressbar.DefaultBytes(res.ContentLength, "正在下载: "+path.Base(target.Address))), res.Body); e != nil {
		logs.Errorf("download failure %s %s %s", name, target.Address, e)
		system.Remove(s.Path(name))
		return e
	} else {
		logs.Infof("download success %s %s %s %s", name, target.Address, logs.Size(n), logs.Cost(begin))
		for _, p := range target.Plugin {
			args := strings.Split(p, " ")
			system.CommandBuild("", args[0], args[1:]...)
		}
		return nil
	}
}
