FROM golang:1.15.0 AS build

ARG BUILD_VERSION=0.0.0-auto
WORKDIR /app

COPY . /app

RUN go build \
    -tags 'osusergo netgo static_build' \
    -ldflags '-X kool-dev/healthz/cmd.version='$BUILD_VERSION' -extldflags "-static"' \
    -o healthz

FROM harborfront/base

COPY --from=build /app/healthz /healthz

ENTRYPOINT [ "/healthz" ]
