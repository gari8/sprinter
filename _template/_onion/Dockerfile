FROM golang:alpine:@@.GoVer@@

ENV GO111MODULE=on

WORKDIR /app
COPY . .
RUN apk add --no-cache git && go get github.com/cespare/reflex
