FROM golang:1.19-alpine

WORKDIR /app

COPY . /app

EXPOSE 8080

RUN go run .