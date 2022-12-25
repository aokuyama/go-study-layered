FROM golang:1.20rc1

RUN apt-get update && \
    go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/stamblerre/gocode@latest && \
    go install -v github.com/rogpeppe/godef@latest

WORKDIR /workspace

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download
