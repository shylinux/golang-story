# {{title "golang-compile"}}
golang是一门简洁高效的编程语言。

- 官网: https://golang.google.cn
- 源码: https://dl.google.com/go/go1.13.5.src.tar.gz
- 开源: https://github.com/golang/go.

## {{chapter "源码解析"}}
下载源码
{{shell "下载源码" "-demo" `wget https://dl.google.com/go/go1.13.5.src.tar.gz`}}

解压源码
{{shell "解压源码" "-demo" `tar xvf go1.13.5.src.tar.gz`}}

项目结构
{{shell "项目结构" "./" `dir go`}}

源码结构
{{shell "源码结构" "./" `dir go/src`}}

编译源码

{{shell "编译源码" "-demo" `cd go/src && ./all.bash`}}

cmd
go
reflect
runtime
syscall

flag
log

命令入口

{{shell "go编译命令" "./go" `sed src/cmd/go/main.go -n -e "1,42p" -e "42a...\n" -e "82,104p" -e "104a...\n"`}}

编译入口

{{shell "go编译命令" "./go" `sed src/cmd/go/internal/work/build.go -n -e "1,42p" -e "42a...\n" -e "82,104p" -e "104a...\n"`}}

编译流程

{{shell "go编译流程" "./go" `sed src/cmd/compile/internal/gc/main.go -n -e "1,42p" -e "42a...\n" -e "132,763p"`}}



