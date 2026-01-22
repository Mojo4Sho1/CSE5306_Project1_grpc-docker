FROM python:3.10-slim

WORKDIR /app

# Copy the Python server code + generated gRPC stubs into the container
COPY . .

# Install only what the Python server needs
RUN pip install --no-cache-dir grpcio protobuf

# gRPC server port
EXPOSE 50051

# Default port (can override later)
ENV PORT=50051

CMD ["python", "echo_server.py"]