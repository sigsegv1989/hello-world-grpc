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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

var (
	countMap map[string]int32 = make(map[string]int32)
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.BatchRequest) (*pb.BatchResponse, error) {
	if len(req.Requests) == 0 {
		log.Print("Request message is empty")
		// Return InvalidArgument if no request messages are provided
		return nil, status.Errorf(codes.InvalidArgument, "No request messages provided")
	}

	log.Printf("Received: %v", req)
	var responses []*pb.Message

	for _, msg := range req.Requests {
		if len(msg.Name) == 0 {
			log.Print("Request's name field is empty")
			// Return InvalidArgument if no request's name are provided
			return nil, status.Errorf(codes.InvalidArgument, "No request's name provided")
		}

		if count, ok := countMap[msg.Name]; ok {
			countMap[msg.Name] = count + 1
		} else {
			countMap[msg.Name] = 1
		}

		response := &pb.Message{
			Name:     msg.Name,
			Greeting: "Hello, " + msg.Name,
			Count:    int32(countMap[msg.Name]),
		}
		responses = append(responses, response)
	}

	batchResponse := &pb.BatchResponse{
		Responses: responses,
	}

	return batchResponse, nil
}

func main() {
	// Define command-line flags with default values
	serverCertPath := flag.String("server-cert", "certs/server/server.crt", "Path to server certificate file (PEM format)")
	serverKeyPath := flag.String("server-key", "certs/server/server.key", "Path to server private key file (PEM format)")
	clientCACertPath := flag.String("client-ca", "certs/ca/ca.crt", "Path to client CA certificate file (PEM format)")
	networkFlag := flag.String("network", "tcp", "Network type (tcp or unix)")
	addressFlag := flag.String("address", ":50051", "Address for the server to listen on")

	// Parse command-line flags
	flag.Parse()

	// Listen on the specified network and address
	listener, err := net.Listen(*networkFlag, *addressFlag)
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
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Printf("Server listening on %s://%s\n", *networkFlag, *addressFlag)
	// Start the gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
