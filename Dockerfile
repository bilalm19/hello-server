FROM golang:1.15.5 as base

WORKDIR /go/bin/src

COPY . .

RUN make test && make build

FROM alpine:latest as prod

WORKDIR /workspace

COPY --from=base /go/bin/src/hello ./hello

EXPOSE 7100

ENTRYPOINT [ "./hello" ]
