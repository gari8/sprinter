FROM golang:alpine

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN go mod init @@.ImportPath@@

RUN apk add --no-cache git && go get github.com/cespare/reflex
