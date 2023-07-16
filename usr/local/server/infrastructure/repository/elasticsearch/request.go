package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

func (s *search) address(str string, arg ...interface{}) string {
	return fmt.Sprintf("http://%s:%d/%s/", s.conf.Host, s.conf.Port, s.conf.Index) + fmt.Sprintf(str, arg...)
}
func (s *search) request(ctx context.Context, method string, url string, data interface{}, res interface{}) error {
	begin := time.Now()
	logs.Infof("search %s request %s %s", method, url, logs.Marshal(data), ctx)
	if err := system.RequestJSON(method, url, data, res); err != nil && err.Error() != "" {
		logs.Warnf("search %s response %s %s %s cost:%s", method, url, err, logs.Marshal(res), logs.Cost(begin), ctx)
		return err
	} else {
		logs.Infof("search %s response %s %s cost:%s", method, url, logs.Marshal(res), logs.Cost(begin), ctx)
		return nil
	}
}
