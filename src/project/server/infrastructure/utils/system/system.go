package system

import (
	"encoding/json"
	"fmt"
	"time"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

func MarshalIndent(v interface{}) string {
	buf, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		logs.Errorf("marshal failure %s", err)
	}
	return string(buf)
}
func Printfln(str string, arg ...interface{}) {
	fmt.Printf(time.Now().Format("2006-01-02T15:04:06.000+0800 "))
	fmt.Printf(str, arg...)
	fmt.Println()
}
