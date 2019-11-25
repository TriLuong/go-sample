FROM golang:latest

WORKDIR /Users/tri/Desktop/project/golang/go-sample

COPY ./ /Users/tri/Desktop/project/golang/go-sample

RUN go mod download

ENTRYPOINT go run main.go