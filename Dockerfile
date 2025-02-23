# Dockerfile
FROM golang:1.24

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4

WORKDIR /app
COPY . .

# If you have a Go program as your action entry point, build it:
RUN go build -o /action ./cmd/main.go

# Start from /app or /github/workspace as needed:
WORKDIR /github/workspace

ENTRYPOINT ["/action"]