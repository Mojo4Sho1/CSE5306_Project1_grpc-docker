import os
import grpc

import echo_pb2
import echo_pb2_grpc


def run():
    server_addr = os.environ.get("SERVER_ADDR", "localhost:50051")

    with grpc.insecure_channel(server_addr) as channel:
        stub = echo_pb2_grpc.EchoServiceStub(channel)

        message = "Hello from the Python client!"
        response = stub.Echo(echo_pb2.EchoRequest(message=message))

        print(f"[Python Client] Sent: {message}")
        print(f"[Python Client] Received original: {response.original}")
        print(f"[Python Client] Received echoed:   {response.echoed}")
        print(f"[Python Client] Received length:   {response.length}")


if __name__ == "__main__":
    run()
