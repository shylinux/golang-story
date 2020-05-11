all:
	@echo && date
	go build -v -o bin/ice.bin src/main.go && chmod u+x bin/ice.bin && chmod u+x bin/ice.sh && ./bin/ice.sh restart
