build:
	go build -o main cmd/web/main.go

help:
	go run cmd/web/main.go --help | go run cmd/tooling/logfmt/main.go

load:
	hey -m GET -c 100 -n 100000 "http://localhost:3000/v1/checks/readiness"

admin:
	go run cmd/tooling/admin/main.go

ready:
	curl -il http://localhost:3000/v1/readiness

live:
	curl -il http://localhost:3000/v1/liveness

# ==============================================================================

BASE_IMAGE_NAME := goecom
VERSION         := "0.0.1-$(shell git rev-parse --short HEAD)"
APP_IMAGE       := $(BASE_IMAGE_NAME):$(VERSION)

# ==============================================================================
# Building containers

all: api

api:
	docker build \
		-f Dockerfile \
		-t $(APP_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Dev envoironment

dev:
	DEV_GOECOM_IMAGE=$(APP_IMAGE) docker compose -f compose.dev.yaml up

swag:
	swag init -g cmd/web/main.go

# ==============================================================================
# Metrics and Tracing

metrics-view-sc:
	expvarmon -ports="localhost:4000" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vuln-check:
	govulncheck ./...