FROM golang:latest

WORKDIR /Users/tri/Desktop/project/golang/go-sample/bin

COPY ./ /Users/tri/Desktop/project/golang/go-sample/bin

RUN go mod download

ENTRYPOINT go run main.go