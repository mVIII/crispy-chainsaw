PROJECT_NAME ?= crispy-chainsaw
SHELL 	 := /bin/bash
executable := crispy-chainsaw
targetOS?= linux
##make clean: deletes all binaries inside bin/
clean:
	@rm -rf bin/*
##make build: generates a binary at bin/$(executable). You can define the target OS you want for the binary by passing targetOS. (e.g. make build targetOS = windows) FYI default targetOS is linux.
build:
	@echo Building $(executable)
	GOOS=$(targetOS) GO111MODULE=on go build -o bin/$(executable) main.go
test:
	go test ./...

.PHONY : help
help : Makefile
	@sed -n 's/^##//p' $<
