# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ENV GO111MODULE=on

WORKDIR /rest-api

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o main .

EXPOSE 8080

CMD [ "/rest-api/main" ]