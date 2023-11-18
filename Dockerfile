FROM golang:1.20.5

RUN apt-get update && \
    go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/stamblerre/gocode@latest && \
    go install -v github.com/rogpeppe/godef@latest

WORKDIR /workspaces

ARG packages="./packages/"
COPY ${packages}domain/go.mod ${packages}domain/
COPY ${packages}domain/go.sum ${packages}domain/

COPY ${packages}app/go.mod ${packages}app/
COPY ${packages}app/go.sum ${packages}app/

COPY ${packages}infra/go.mod ${packages}infra/
COPY ${packages}infra/go.sum ${packages}infra/

ARG apps="./"
COPY ${apps}cui/go.mod ${apps}cui/
COPY ${apps}cui/go.sum ${apps}cui/

COPY ${apps}http_server/go.mod ${apps}http_server/
COPY ${apps}http_server/go.sum ${apps}http_server/

COPY ${apps}go.work ${apps}
COPY ${apps}go.work.sum ${apps}

RUN go work sync
