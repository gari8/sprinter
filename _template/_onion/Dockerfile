FROM golang:1.15-alpine

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
RUN apk add --no-cache alpine-sdk git && go get github.com/cespare/reflex