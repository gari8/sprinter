FROM golang:1.16beta1-alpine3.12

ENV GO111MODULE=on

WORKDIR /app
COPY . .
RUN apk add --no-cache git && go get github.com/cespare/reflex
