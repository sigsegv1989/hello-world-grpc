package main

import (
	"context"
	"log"

	pb "github.com/sigsegv1989/hello-world-grpc/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	message := &pb.Message{
		Name:     "Alice",
		Greeting: "Hi there!",
		Count:    0,
	}
	response, err := c.SayHello(context.Background(), message)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	log.Printf("Response: %v", response)
}
