FROM golang:1.23.3-alpine3.20 AS build
ENV CGO_ENABLED=0
ARG BUILD_REF

# Copy the source code into the container.
COPY . /app

# Build the admin binary.
WORKDIR /app/cmd/tooling/admin
RUN go build -ldflags "-X main.build=${BUILD_REF}"

WORKDIR /app/cmd/tooling/logfmt
RUN go build

# Build the service binary.
WORKDIR /app/cmd/web
RUN go build -ldflags "-X main.build=${BUILD_REF}"

FROM golang:1.23.3-alpine3.20 AS develop
ARG BUILD_DATE
ARG BUILD_REF
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["sh", "-c", "go run cmd/tooling/admin/main.go migrate-seed && air -c .air.toml | go run cmd/tooling/logfmt/main.go"]

FROM alpine3.20 AS deploy
ARG BUILD_DATE
ARG BUILD_REF
RUN addgroup -g 1000 -S goecom && \
    adduser -u 1000 -h /app -G goecom -S goecom
COPY --from=build --chown=goecom:goecom /app/cmd/tooling/admin/admin /app/admin
COPY --from=build --chown=goecom:goecom /app/cmd/tooling/logfmt/logfmt /app/logfmt
COPY --from=build --chown=goecom:goecom /app/cmd/web/web /app/web
WORKDIR /app
USER goecom
CMD ["sh", "-c", "./admin migrate-seed && ./web | ./logfmt"]