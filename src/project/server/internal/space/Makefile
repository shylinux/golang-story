all:

idl:
	protoc --go_out=./idl --go-grpc_out=./idl ./idl/*.proto
test:
	go test -v --count=1 ./idl/test
.PHONY: idl
