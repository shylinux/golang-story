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
func PwdModPath() string {
	file, _ := os.Getwd()
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
func Size(size int64) string {
	if size > 1<<30 {
		return fmt.Sprintf("%.2fG", float64(size)/(1<<30))
	} else if size > 1<<20 {
		return fmt.Sprintf("%.2fM", float64(size)/(1<<20))
	} else if size > 1<<10 {
		return fmt.Sprintf("%.2fK", float64(size)/(1<<10))
	} else {
		return fmt.Sprintf("%sB", size)
	}
}
func Cost(begin time.Time) string {
	return fmt.Sprintf("%.2fms", float64(time.Now().Sub(begin))/float64(time.Millisecond))
}
func Percent(size int64, total int64) string {
	if total == 0 {
		total = 2*size + 1
	}
	return fmt.Sprintf("%d%%", int(size*100/total))
}
func Marshal(v interface{}) string {
	buf, _ := json.Marshal(v)
	return string(buf)
}
func MarshalIndent(v interface{}) string {
	buf, _ := json.MarshalIndent(v, "", "  ")
	return string(buf)
}
