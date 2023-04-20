FROM golang:1.19.3-alpine3.16

ENV CGO_ENABLED=1

RUN apk update
RUN apk upgrade
RUN apk add --update figlet
RUN apk add build-base

RUN go install github.com/cespare/reflex@latest
