FROM golang:1.22-bookworm

COPY . /app

WORKDIR /app

RUN go build main.go

FROM debian:bookworm

COPY --from=0 /app/main /app/main

WORKDIR /app
