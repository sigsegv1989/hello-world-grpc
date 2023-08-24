package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/sigsegv1989/hello-world-grpc/api/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	countMap map[string]int32 = make(map[string]int32)
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Received: %v", in)
	if count, ok := countMap[in.Name]; ok {
		countMap[in.Name] = count + 1
	} else {
		countMap[in.Name] = 1
	}
	return &pb.Message{
		Greeting: "Hello, " + in.Name,
		Count:    int32(countMap[in.Name]),
	}, nil
}

func main() {
	// Define command-line flags with default values
	serverCertPath := flag.String("server-cert", "certs/server/server.crt", "Path to server certificate file (PEM format)")
	serverKeyPath := flag.String("server-key", "certs/server/server.key", "Path to server private key file (PEM format)")
	clientCACertPath := flag.String("client-ca", "certs/ca/ca.crt", "Path to client CA certificate file (PEM format)")
	flag.Parse()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Load server certificate and key
	serverCert, err := tls.LoadX509KeyPair(*serverCertPath, *serverKeyPath)
	if err != nil {
		log.Fatalf("Failed to load server certificate: %v", err)
	}

	// Load trusted client CA certificates
	clientCACert, err := ioutil.ReadFile(*clientCACertPath)
	if err != nil {
		log.Fatalf("Failed to read client CA certificate: %v", err)
	}
	clientCACertPool := x509.NewCertPool()
	clientCACertPool.AppendCertsFromPEM(clientCACert)

	// Create a credentials object using the server certificate and client CA
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    clientCACertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	// Set up gRPC server with the credentials
	s := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterGreeterServer(s, &server{})
	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
