# ---- Build stage ----
FROM golang:1.25.6 AS builder

WORKDIR /app

# Copy the whole Go module so imports work (pb + server + go.mod)
COPY go/ ./go/

WORKDIR /app/go/server

RUN go mod download
RUN go build -o server .

# ---- Run stage ----
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/go/server/server ./server

EXPOSE 50051
ENV PORT=50051

CMD ["./server"]