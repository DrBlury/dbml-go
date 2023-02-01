# Build image using go 1.19
FROM golang:1.19-alpine AS build
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the dbml-gen-go-model binary
RUN go build ./cmd/dbml-gen-go-model/main.go
# Copy the binary to the production image from the builder stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main ./dbml-go/-generator
