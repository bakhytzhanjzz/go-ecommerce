# Build stage
FROM golang:1.21-alpine AS builder

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create and change to the app directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

# Install dependencies (if needed)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose the port the app runs on
EXPOSE 8001

# Command to run the application
CMD ["./main"]