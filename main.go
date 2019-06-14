package main

import (
	"fmt"

	"github.com/po3rin/dockerdot/docker2dot"
)

func main() {
	dockerfile := []byte(`
	FROM golang:1.12 as builder

	WORKDIR /go/wedding

	COPY . .

	ENV GO111MODULE=on
	RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


	FROM alpine:latest

	RUN apk --no-cache add ca-certificates

	WORKDIR /app
	COPY --from=builder /go/wedding .
	COPY ./client .
	`)
	dot, err := docker2dot.Docker2Dot(dockerfile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dot))
}
