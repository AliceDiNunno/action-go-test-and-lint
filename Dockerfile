# Dockerfile
FROM golang:1.24

# Install golangci-lint
# (Pick a version or use "latest"—here we use v1.52.2 as example)
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2

WORKDIR /app
COPY . .

# If you have a Go program as your action entry point, build it:
RUN go build -o /action ./cmd/main.go

# Start from /app or /github/workspace as needed:
WORKDIR /github/workspace

ENTRYPOINT ["/action"]