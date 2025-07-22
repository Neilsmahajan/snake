# Build stage
FROM golang:1.24.5-alpine AS builder

# Set working directory
WORKDIR /app

# Install git for dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o snake ./cmd/

# Final stage
FROM alpine:latest

# Install necessary packages for terminal functionality
RUN apk --no-cache add ca-certificates ncurses

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/snake .

# Set terminal type for better compatibility
ENV TERM=xterm-256color

# Run the application
CMD ["./snake"]
