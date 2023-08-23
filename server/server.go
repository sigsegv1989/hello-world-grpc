package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sigsegv1989/hello-world-grpc/api" // Import the shared package
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Received: %v", in)
	return &pb.Message{
		Greeting: "Hello, " + in.Name,
		Count:    in.Count + 1,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
