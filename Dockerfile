FROM golang:1.14.2-alpine3.11

ENV GO111MODULE=on

WORKDIR /doc-server
COPY go.mod .

RUN go mod tidy
COPY . .