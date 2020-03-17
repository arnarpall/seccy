FROM golang:1.14 AS build
ARG ld_flags

WORKDIR /build
ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install \
    -a \
    -trimpath \
    #-ldflags ${ld_flags} \
    ./...

FROM alpine:latest
COPY --from=build /go/bin/seccy-service ./
ENTRYPOINT ["./seccy-service"]