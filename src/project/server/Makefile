all:
	go build -v -o bin/demo cmd/main.go && kill -INT `cat log/service.pid`
idl:
	protoc --go_out=./idl --go-grpc_out=./idl ./idl/*.proto
	# protoc --js_out=import_style=commonjs:./usr/vue-element-admin/src/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./usr/vue-element-admin/src/ ./idl/*.proto
test:
	go test -v --count=1 ./idl/test
pack:
	tar zcvf service.tar.gz bin config usr/vue-element-admin/dist
clean:
	rm -f bin/* log/* idl/test/log/*
.PHONY: idl test pack clean
