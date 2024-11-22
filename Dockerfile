# Dockerfile for building and running the Prometheus metrics service

# Stage 1: Build the Go binary using a Go image
FROM golang:1.18 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o /fake-metrics

# Stage 2: Create a small image using Alpine and copy the binary from the builder
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the builder
COPY --from=builder /fake-metrics .

# Command to run the executable
CMD ["./fake-metrics"]

# Expose port 9000 for metrics
EXPOSE 9000
