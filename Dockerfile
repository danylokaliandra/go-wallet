FROM golang:1.17-alpine

LABEL maintainer="Danylo Kaliandra <danylokaliandra@correo.ugr.es>"

ENV TEST_DIR=/app/test/

RUN addgroup -S mywallet && adduser -S mywallet -G mywallet

USER mywallet

WORKDIR $TEST_DIR

COPY go.mod /app/

RUN go mod download

RUN go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT ["task", "test"]