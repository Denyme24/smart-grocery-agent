# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o smart-grocery-agent ./cmd/main.go

# Run stage
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/smart-grocery-agent .

# Copy .env file
COPY --from=builder /app/.env .

# Expose the application port
EXPOSE 3000

# Command to run the executable
CMD ["./smart-grocery-agent"]