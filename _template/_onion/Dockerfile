FROM golang:1.14.2-alpine3.11

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
RUN apk add --no-cache alpine-sdk git && go get github.com/cespare/reflex