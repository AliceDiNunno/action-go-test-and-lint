FROM golang:1.24.0 AS builder

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build our custom action binary
RUN CGO_ENABLED=0 go build -o /action ./cmd/main.go

FROM alpine:3.18

RUN apk add --no-cache ca-certificates

COPY --from=builder /action /action
COPY --from=builder /go/bin/golangci-lint /usr/local/bin/golangci-lint

WORKDIR /github/workspace

ENTRYPOINT ["/action"]