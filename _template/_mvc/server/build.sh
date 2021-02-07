#!/bin/bash

go mod tidy
reflex -r '(\.go$|go\.mod)' -s go run main.go
