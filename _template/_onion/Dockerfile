FROM golang:1.15.6-alpine

ENV GO111MODULE=on

WORKDIR /app
COPY . .
RUN apk add --no-cache git && go get github.com/cespare/reflex
