#!/bin/bash

go tool dist list \
  | awk '{split($0,a,"/"); \
  printf "env GOOS=%s GOARCH=%s go build -o bin/%s/%s/pm cmd/pm/main.go\n",a[1], a[2], a[1], a[2]}' \
  | sh
