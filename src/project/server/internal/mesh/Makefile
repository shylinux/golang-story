export CGO_ENABLED=1
def: matrix idl server restart
all: matrix idl server deploy restart test pack
matrix:
	go build -v -o bin/matrix cmd/cmds.go
idl:
	./bin/matrix unpack protoc
	./bin/matrix generate
server:
	go build -v -o bin/server cmd/main.go
deploy:
	./bin/matrix deploy node java git golang
	./bin/matrix deploy nginx redis consul pulsar es
restart:
	./bin/matrix server restart
test:
	go test -v --count=1 ./idl/ts
	go test -v --count=1 ./idl/test
pack:
	tar zcvf service.tar.gz bin config usr/vue-element-admin/dist
clean:
	rm -rf bin/* log/* idl/test/log/* usr/vue-element-admin/dist/*
.PHONY: matrix idl server restart test pack clean
