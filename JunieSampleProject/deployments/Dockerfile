# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application for multiple platforms
RUN mkdir -p /app/build && \
    GOOS=linux GOARCH=amd64 go build -o /app/build/myapp-linux-amd64 ./cmd/myapp && \
    GOOS=windows GOARCH=amd64 go build -o /app/build/myapp-windows-amd64.exe ./cmd/myapp && \
    GOOS=darwin GOARCH=amd64 go build -o /app/build/myapp-darwin-amd64 ./cmd/myapp

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/build/myapp-linux-amd64 /app/myapp

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["/app/myapp"]