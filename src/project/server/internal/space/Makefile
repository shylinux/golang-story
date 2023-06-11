all:
	go build -v -o bin/space cmd/main.go && kill -INT `cat log/service.pid`
idl:
	protoc --go_out=./idl --go-grpc_out=./idl ./idl/*.proto
	go run cmd/generate.go
test:
	go test -v --count=1 ./idl/test
clean:
	rm -f bin/* log/* idl/test/log/*
.PHONY: idl test clean
