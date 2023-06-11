package logs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"go.opentelemetry.io/otel/trace"
)

func TraceID(ctx context.Context) string {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return fmt.Sprintf("%s-%s", span.TraceID().String(), span.SpanID())
	}
	return ""
}
func ModPath(skip int) string {
	_, file, _, _ := runtime.Caller(skip)
	ls := strings.Split(file, "/")
	for i := len(ls) - 1; i > 0; i-- {
		p := path.Join(strings.Join(ls[:i], "/"), "go.mod")
		if _, e := os.Stat(p); e == nil {
			buf, _ := ioutil.ReadFile(p)
			head := bytes.SplitN(buf, []byte("\n"), 2)
			return path.Join(strings.Split(string(head[0]), " ")[1], path.Join(ls[i:]...))
		}
	}
	return file
}
func FileLine(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	list := strings.Split(file, "/")
	if len(list) > 2 {
		list = list[len(list)-2:]
	}
	return fmt.Sprintf("%s:%d", path.Join(list[:]...), line)
}
func FuncName(skip int) string {
	fun, _, _, _ := runtime.Caller(skip)
	return path.Base(runtime.FuncForPC(fun).Name())
}
func Cost(begin time.Time) string {
	return fmt.Sprintf("%.2fms", float64(time.Now().Sub(begin))/float64(time.Millisecond))
}
func Marshal(v interface{}) string {
	buf, _ := json.Marshal(v)
	return string(buf)
}
