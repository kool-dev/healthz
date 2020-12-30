FROM golang:1.15.0 AS build

WORKDIR /app

COPY . /app

RUN go build -tags 'osusergo netgo static_build' -ldflags '-extldflags "-static"' -o healthz

FROM harborfront/base

COPY --from=build /app/healthz /healthz

ENTRYPOINT [ "/healthz" ]
