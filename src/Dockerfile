FROM golang:1.17 as build-env

ENV APPDIR=$GOPATH/src/github.com/hyorimitsu/hello-cloud-profiler/src \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR $APPDIR
COPY . $APPDIR/

RUN go mod download
RUN go mod verify

RUN go build -ldflags "-s -w" -o /app $APPDIR/cmd/api/main.go

FROM gcr.io/distroless/static

COPY --from=build-env /app /

CMD ["/app"]
