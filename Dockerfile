# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o fundraiserApplication ./cmd/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /app/fundraiserApplication ./fundraiserApplication

EXPOSE 8080

USER nonroot:nonroot

# Use below entry point to create admin
# ENTRYPOINT ["./fundraiserApplication", "create-admin", "admin", "password"]
ENTRYPOINT ["./fundraiserApplication"]