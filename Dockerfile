FROM golang:1.20rc1

RUN apt-get update && \
    go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/stamblerre/gocode@latest && \
    go install -v github.com/rogpeppe/godef@latest

WORKDIR /workspace

COPY ./domain/go.mod ./domain/
COPY ./domain/go.sum ./domain/

COPY ./app/go.mod ./app/
COPY ./app/go.sum ./app/

COPY ./infra/go.mod ./infra/
COPY ./infra/go.sum ./infra/

COPY ./cui/go.mod ./cui/
COPY ./cui/go.sum ./cui/

COPY ./http_server/go.mod ./http_server/
COPY ./http_server/go.sum ./http_server/

COPY ./go.work ./
COPY ./go.work.sum ./

RUN go work sync
