FROM golang:1.4.2
MAINTAINER Brainly <open.source@brainly.com>

### Install necessary libraries ###
RUN apt-get update && apt-get install -y \
    tesseract-ocr \
    tesseract-ocr-pol

### Install godep, dependency manager for golang ###
RUN go get github.com/tools/godep

### Set up working directory ###
RUN mkdir -p /go/src/github.com/brainly/go-tesseract
WORKDIR /go/src/github.com/brainly/go-tesseract
ADD . /go/src/github.com/brainly/go-tesseract

### Build the project including dependencies ###
# RUN godep go install

### Make sure the code passes the linter ###
# RUN godep golint

### Make sure code is vetted ###
# RUN godep go vet

### Make sure all tests pass ###
# RUN godep go test ./...

# ENTRYPOINT ["/go/bin/go-tesseract"]
