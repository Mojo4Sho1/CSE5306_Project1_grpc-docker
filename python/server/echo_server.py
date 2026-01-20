from concurrent import futures
import os

import grpc

import echo_pb2
import echo_pb2_grpc


class EchoService(echo_pb2_grpc.EchoServiceServicer):
    def Echo(self, request, context):
        msg = request.message
        reply = echo_pb2.EchoReply(
            original=msg,
            echoed=msg.upper(),
            length=len(msg),
        )
        return reply


def serve():
    port = os.environ.get("PORT", "50051")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    echo_pb2_grpc.add_EchoServiceServicer_to_server(EchoService(), server)

    # IMPORTANT: bind to all interfaces (needed later for Docker)
    server.add_insecure_port(f"0.0.0.0:{port}")

    print(f"[Python Server] Listening on 0.0.0.0:{port}")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
