all:
	protoc --go_out=./idl --go-grpc_out=./idl ./idl/*.proto
	go build -v -o bin/space main.go
test:
	go test -v --count=1 ./idl/test


