# syntax=docker/dockerfile:1

FROM golang:1-alpine

WORKDIR /app

COPY go.mod ./



LABEL "com.example.vendor"="ACME Incorporated"
LABEL com.example.label-with-value="foo"
LABEL version="1.0"
LABEL description="This text illustrates"
LABEL alias="yourtool"


MAINTAINER "gleb mojeico"
MAINTAINER "gleb mojeico"



RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]