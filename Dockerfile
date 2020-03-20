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

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO /bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

FROM alpine:latest
COPY --from=build /go/bin/seccy-service ./
COPY --from=build /bin/grpc_health_probe ./

ENTRYPOINT ["./seccy-service"]