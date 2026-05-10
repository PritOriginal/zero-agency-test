# Use the official Golang image as the base
FROM golang:1.26-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory inside the container
WORKDIR /build

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application source
COPY . .

# Build the Go binary
RUN go build -o /bot-openai ./cmd/bot-openai

# Final lightweight stage
FROM alpine:3.21 AS final

# Copy the compiled binary from the builder stage
COPY --from=builder /bot-openai /bin/bot-openai

COPY ./configs/config.yaml /bin/configs/config.yaml

# Run the application
CMD ["/bin/bot-openai", "--config=/bin/configs/config.yaml"]