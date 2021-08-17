#!/bin/bash

go mod init @@.ImportPath@@
go get github.com/gari8/sprinter
go get github.com/go-chi/chi
go mod tidy