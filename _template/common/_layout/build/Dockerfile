FROM golang:alpine

ENV GO113MODULE=on

WORKDIR /app

COPY ../@@.ImportPath@@ .

RUN go mod tidy

RUN apk update && apk add --no-cache git && go get github.com/cespare/reflex
