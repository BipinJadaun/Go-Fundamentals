# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY /src ./

RUN go build -o /go-fundamentals

EXPOSE 8080 8080

ENTRYPOINT [ "/go-fundamentals/src/main.go" ]