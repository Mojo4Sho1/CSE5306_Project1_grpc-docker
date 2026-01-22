FROM python:3.10-slim

WORKDIR /app

# Copy the Python client code + generated gRPC stubs into the container
COPY . .

# Install what the client needs
RUN pip install --no-cache-dir grpcio protobuf

# Default server address (we will override this when running containers)
ENV SERVER_ADDR=localhost:50051

CMD ["python", "echo_client.py"]
