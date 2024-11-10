# Build stage
FROM golang:1.21-alpine AS builder
ARG PB_VERSION=0.22.21

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy and download dependencies first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 go build -o pb ./main.go

# Final stage
FROM alpine:latest

# Install required runtime dependencies
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/pb /app/pb

# Create directory for persistent data
RUN mkdir -p /app/pb_data

EXPOSE 8090

CMD ["/app/pb", "serve", "--http=0.0.0.0:8090"]
