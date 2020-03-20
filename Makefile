VERSION ?= $(shell git describe --exact-match 2> /dev/null || \
           git describe --match=$(git rev-parse --short=8 HEAD) --always --dirty --abbrev=8 || \
           echo "dev")
BUILD ?= $(shell git rev-parse --short HEAD)
REPOSITORY := arnar.io/seccy
BUILD_DATE = $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LD_FLAGS := "-X github.com/arnarpall/seccy/internal/version.BuildVersion=${VERSION} -X github.com/arnarpall/seccy/internal/version.BuildDate=${BUILD_DATE}"

build:
	go build -ldflags ${LD_FLAGS} ./...

service:
	cd cmd/seccy-service/ && go build -ldflags ${LD_FLAGS}

run-service: service
	cmd/seccy-service/seccy-service

run-service-dev: service
	cmd/seccy-service/seccy-service --encryption-key test --store-path /tmp/test.vault

rest-api:
	cd cmd/seccy-rest-api/ && go build -ldflags ${LD_FLAGS}

run-rest-api: rest-api
	cmd/seccy-rest-api/seccy-rest-api

proto:
	- protoc -I api/proto ./api/proto/seccy/seccy.proto --go_out=plugins=grpc:.

cli:
	cd cmd/seccy-cli/ && go build -ldflags ${LD_FLAGS}

install:
	cd cmd/seccy-cli/ && go install -ldflags ${LD_FLAGS}

test:
	go test ./...

.PHONY: image
image:  ## Build the docker image locally
	docker build --build-arg ld_flags=${LD_FLAGS} -t ${REPOSITORY}:${VERSION} .
