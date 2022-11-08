FROM golang:1.17 as builder

#
RUN mkdir -p $GOPATH/src/github.com/Mirobidjon/reverse-proxy
WORKDIR $GOPATH/src/github.com/Mirobidjon/reverse-proxy

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    # go mod vendor && \
    make build && \
    mv ./bin/reverse-proxy / && \
    mv proxy.yaml /

FROM alpine
COPY --from=builder reverse-proxy .
RUN mkdir /etc/reverse-proxy/
COPY --from=builder proxy.yaml /etc/reverse-proxy/

ENTRYPOINT ["/reverse-proxy"]
