{{title "golang-runtime"}}
{{brief "语言简介" `golang是一门简洁高效的编程语言。`}}

{{refer "官方网站" `
官网: https://golang.google.cn
源码: https://dl.google.com/go/go1.13.5.src.tar.gz
开源: https://github.com/golang/go.
`}}

{{order "官方包" `
src/internal
src/builtin
src/runtime
src/context
src/sync
`}}

{{chapter "程序入口"}}
{{shell "查看源码" "usr" `cat > hi.go <<END
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
END`}}

{{shell "编译代码" "usr" `go build -o hi hi.go`}}
{{shell "执行程序" "usr" `./hi`}}
{{shell "汇编程序" "usr" `objdump -d hi > hi.s`}}
{{shell "程序信息" "usr" `readelf -h hi`}}

{{shell "程序入口" "usr" `sed -n "/<_rt0_amd64_linux>:/,+2p" hi.s`}}
{{shell "程序入口" "usr" `sed -n "/<_rt0_amd64>:/,+85p" hi.s`}}
{{shell "程序入口" "usr" `cat go/src/runtime/rt0_linux_amd64.s`}}
{{shell "程序入口" "usr" `head -n241 go/src/runtime/asm_amd64.s`}}

{{chapter "调度模型"}}

{{order "相关代码" `
runtime/proc.go
runtime/runtime2.go
runtime/stack.go
runtime/debug.go
runtime/extern.go
`}}

{{chapter "异常处理"}}

{{order "相关代码" `
runtime/panic.go
`}}

{{chapter "协程并发"}}

{{order "相关代码" `
runtime/rwmutex.go
runtime/sema.go
sync/rwmutex.go
sync/mutex.go
sync/cond.go
sync/map.go
sync/pool.go
sync/once.go
sync/waitgroup.go
`}}

{{chapter "协程通信"}}

{{order "相关代码" `
runtime/chan.go
runtime/select.go
context/context.go
`}}

{{chapter "内存分配"}}

{{order "相关代码" `
runtime/malloc.go
`}}

{{chapter "垃圾回收"}}

{{order "相关代码" `
runtime/mgc.go
`}}

{{chapter "数据类型"}}

{{order "相关代码" `
builtin/builtin.go
runtime/type.go
runtime/utf8.go
runtime/string.go
runtime/complex.go
runtime/float.go
`}}

{{shell "基本类型" "usr" `sed -n "/^type /p" go/src/builtin/builtin.go`}}
{{shell "基本函数" "usr" `sed -n "/^func /p" go/src/builtin/builtin.go`}}

{{chapter "数据结构"}}

{{order "相关代码" `
runtime/alg.go
runtime/map.go
runtime/slice.go
`}}

{{section "slice"}}

{{shell "数据类型" "usr" `sed -n "/type slice/,+4p" go/src/runtime/slice.go`}}
{{shell "构造数据" "usr" `sed -n "/func makeslice(/,+16p" go/src/runtime/slice.go`}}
{{shell "追加数据" "usr" `sed -n "/func growslice(/,+38p" go/src/runtime/slice.go`}}
