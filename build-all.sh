#!/bin/bash

#Build for unixy OSes
go tool dist list \
  | awk '{split($0,a,"/"); \
  printf "env GOOS=%s GOARCH=%s go build -o bin/%s/%s/pm cmd/pm/main.go\n",a[1], a[2], a[1], a[2]}' \
  | grep -v "GOOS=windows" | sh

#Build for windows with .exe extension
go tool dist list \
  | awk '{split($0,a,"/"); \
  printf "env GOOS=%s GOARCH=%s go build -o bin/%s/%s/pm.exe cmd/pm/main.go\n",a[1], a[2], a[1], a[2]}' \
  | grep "GOOS=windows" | sh