all:
	go build -v -o bin/matrix cmd/cmds.go
	go build -v -o bin/server cmd/main.go && kill -INT `cat log/service.pid`
idl:
	protoc --go_out=./idl --go-grpc_out=./idl ./idl/*.proto
	go run cmd/generate.go
	# protoc --govalidators_out=./idl --go_out=./idl --go-grpc_out=./idl ./idl/*.proto
	# protoc --js_out=import_style=commonjs:./usr/vue-element-admin/src/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./usr/vue-element-admin/src/ ./idl/*.proto
test:
	go test -v --count=1 ./idl/test
tool:
	# go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
pack:
	tar zcvf service.tar.gz bin config usr/vue-element-admin/dist
clean:
	rm -rf bin/* log/* idl/test/log/* usr/vue-element-admin/dist/*
.PHONY: idl test pack clean
