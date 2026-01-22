# CSE5306 Project 1 — gRPC + Docker (Python + Go)

This project implements a simple **gRPC Echo service** in two languages (**Python** and **Go**) using the **same proto file**.  
Both clients can communicate with both servers using **cross-language communication**.

## Folder Structure

CSE5306_Project1_grpc-docker/
    proto/
        echo.proto

    python/
        server/
            Dockerfile
            echo_server.py
            echo_pb2.py
            echo_pb2_grpc.py
        client/
            Dockerfile
            echo_client.py
            echo_pb2.py
            echo_pb2_grpc.py

    go/
        go.mod
        go.sum
        pb/
            echo.pb.go
            echo_grpc.pb.go
        server/
            Dockerfile
            main.go
        client/
            Dockerfile
            main.go

## How to Compile and Run (Docker)

All commands below are run from the **repo root**.

### Build Images

#### Python Server Image
docker build -t python-echo-server -f python/server/Dockerfile python/server

#### Python Client Image
docker build -t python-echo-client -f python/client/Dockerfile python/client

#### Go Server Image
docker build -t go-echo-server -f go/server/Dockerfile .

#### Go Client Image
docker build -t go-echo-client -f go/client/Dockerfile .

### Create Docker Network
docker network create grpc-net

### Cross-language Runs

#### Video 1: Python client → Go server

Terminal 1 (Go server):
docker run --rm -it --name go-server --network grpc-net -p 50052:50051 go-echo-server

Terminal 2 (Python client):
docker run --rm --network grpc-net -e SERVER_ADDR=go-server:50051 python-echo-client

#### Video 2: Go client → Python server

Terminal 1 (Python server):
docker run --rm -it --name python-server --network grpc-net -p 50051:50051 python-echo-server

Terminal 2 (Go client):
docker run --rm --network grpc-net -e SERVER_ADDR=python-server:50051 go-echo-client

## Local Runs (Optional)

### Python server
python python/server/echo_server.py

### Python client
python python/client/echo_client.py

### Go server
cd go/server
go run .

### Go client
cd go/client
go run .

## Anything Unusual / Notes

- A custom Docker network (`grpc-net`) is used so containers can communicate by container name.
- Both servers listen on port `50051` inside the container.
- The Go Docker builds use a multi-stage build (build stage + minimal runtime stage).

## External Sources Referenced

- Docker documentation (Get Started / networking concepts)
- gRPC documentation (Python + Go examples)
- Protocol Buffers / protoc documentation
- ChatGPT 5.2 Thinking
