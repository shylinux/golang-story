{{title "golang-compile"}}
{{brief "语言简介" `golang是一门简洁高效的编程语言。`}}

{{refer "官方网站" `
官网: https://golang.google.cn
源码: https://dl.google.com/go/go1.13.5.src.tar.gz
开源: https://github.com/golang/go.
`}}

{{order "官方包" `
src/cmd
src/go
src/reflect
src/runtime
src/syscall
src/unsafe
src/plugin
src/debug
`}}

{{chapter "下载源码"}}
{{shell "下载源码" "usr" "install" `wget https://dl.google.com/go/go1.13.5.src.tar.gz`}}
{{shell "解压源码" "usr" "install" `tar xvf go1.13.5.src.tar.gz`}}

{{shell "项目结构" "usr" `dir go`}}
{{shell "源码结构" "usr" `dir go/src`}}
{{shell "编译源码" "usr" "compile" `cd go/src && ./all.bash`}}

{{chapter "命令入口"}}
{{shell "命令入口" "usr/go" `sed src/cmd/go/main.go -n -e "1,42p" -e "42a...\n" -e "82,104p" -e "104a...\n"`}}
{{shell "构建入口" "usr/go" `sed src/cmd/go/internal/work/build.go -n -e "1,42p" -e "42a...\n" -e "82,104p" -e "104a...\n"`}}
{{shell "编译入口" "usr/go" `sed src/cmd/compile/internal/gc/main.go -n -e "1,42p" -e "42a...\n" -e "132,763p"`}}

{{chapter "闭包"}}
cmd/compile/internal/gc/closure.go

{{chapter "遍历"}}
cmd/compile/internal/gc/range.go

