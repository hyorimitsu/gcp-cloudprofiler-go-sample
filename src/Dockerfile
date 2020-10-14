FROM golang:1.14.2-alpine3.11 AS build-env

ADD . /go/src/github.com/hyorimitsu/gcp-cloudprofiler-go-sample/src
WORKDIR /go/src/github.com/hyorimitsu/gcp-cloudprofiler-go-sample/src

RUN go build .

FROM alpine:3.11
COPY --from=build-env /go/src/github.com/hyorimitsu/gcp-cloudprofiler-go-sample/src/src /src
ENV GOPATH=/go
ENV APP_VERSION="1.0.0"

ENTRYPOINT [ "./src","--host", "0.0.0.0", "--port", "80", "--scheme", "http"]
