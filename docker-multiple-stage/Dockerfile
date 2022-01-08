# https://tutorialedge.net/golang/go-multi-stage-docker-tutorial/

# Multi-Stage Dockerfile

# 1st stage - compile go program
FROM golang:1.12.0-alpine3.9 AS builder
RUN apk add git
RUN mkdir /go/src/go-app
ADD . /go/src/go-app
WORKDIR /go/src/go-app
RUN export GOBIN=$GOPATH/bin
RUN go get ./...
RUN go build -o main .

# 2nd stage - run compiled program
FROM alpine:3.9 AS production
COPY --from=builder /go/src/go-app .
CMD ["./main"]

# $ docker build -t go-multi-stage .

# $ docker images
# REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
# go-multi-stage          latest              12dd51472827        24 seconds ago      13.2MB

# docker run -d -p 8080:8080 go-multi-stage