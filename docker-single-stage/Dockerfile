# https://tutorialedge.net/golang/go-multi-stage-docker-tutorial/

# Single-Stage Dockerfile

# Specify the base image for go application
FROM golang:1.12.0-alpine3.9

# Install git to use go get
RUN apk add git

# Create an /app directory within image 
RUN mkdir /go/src/go-app

# Copy source file in the current directory into /app
ADD . /go/src/go-app

# Specify that any futher commands will execute inside /app
WORKDIR /go/src/go-app

# Specify GOBIN for go packages
RUN export GOBIN=$GOPATH/bin

# Download packages
RUN go get ./...

# Run 'go build' to compile the binary executable of go application
RUN go build -o main .

# Run compiled go application
CMD ["/go/src/go-app/main"]

# $ docker build -t go-single-stage .

# $ docker images
# REPOSITORY                    TAG                 IMAGE ID            CREATED             SIZE
# go-single-stage               latest              761b9dd5f9a4        4 seconds ago       380MB

# docker run -d -p 8080:8080 go-single-stage