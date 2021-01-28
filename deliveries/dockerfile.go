FROM golang:latest AS builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/appa
COPY go.mod .
RUN go mod download
COPY . . 
RUN go install

CMD ["cmd/main"]