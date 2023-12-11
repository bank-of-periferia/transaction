FROM golang:1.16-alpine AS builder
ENV PATH="$PATH:/bin/bash" \
    BENTO4_BIN="/opt/bento4/bin" \
    PATH="$PATH:/opt/bento4/bin"


WORKDIR /go/src

ENTRYPOINT [ "executable" ]