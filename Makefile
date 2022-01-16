SHELL := /bin/bash

build:
	go build -o bin/pm cmd/pm/main.go

build-all:
	./build-all.sh