# ---- Build stage ----
FROM golang:1.25.6 AS builder

WORKDIR /app

# Copy the whole Go module so imports work (pb + client + go.mod)
COPY go/ ./go/

WORKDIR /app/go/client

RUN go mod download
RUN go build -o client .

# ---- Run stage ----
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/go/client/client ./client

# Default server address (we will override when we run)
ENV SERVER_ADDR=localhost:50051

CMD ["./client"]