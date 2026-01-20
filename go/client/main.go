package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Mojo4Sho1/CSE5306_Project1_grpc-docker/go/pb"
)

func main() {
	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = "localhost:50051"
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	msg := "Hello from the Go client!"
	resp, err := client.Echo(ctx, &pb.EchoRequest{Message: msg})
	if err != nil {
		panic(err)
	}

	fmt.Printf("[Go Client] Sent: %s\n", msg)
	fmt.Printf("[Go Client] Received original: %s\n", resp.GetOriginal())
	fmt.Printf("[Go Client] Received echoed:   %s\n", resp.GetEchoed())
	fmt.Printf("[Go Client] Received length:   %d\n", resp.GetLength())
}
