package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"

	pb "github.com/Mojo4Sho1/CSE5306_Project1_grpc-docker/go/pb"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
}

func (s *echoServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoReply, error) {
	msg := req.GetMessage()
	return &pb.EchoReply{
		Original: msg,
		Echoed:   strings.ToUpper(msg),
		Length:  int32(len(msg)),
	}, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterEchoServiceServer(server, &echoServer{})

	fmt.Printf("[Go Server] Listening on 0.0.0.0:%s\n", port)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
