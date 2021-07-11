FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git && apk add build-base

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download
