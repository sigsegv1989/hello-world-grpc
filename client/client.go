package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"

	pb "github.com/sigsegv1989/hello-world-grpc/api/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// Define command-line flags with default values
	clientCertPath := flag.String("client-cert", "certs/client/client.crt", "Path to client certificate file (PEM format)")
	clientKeyPath := flag.String("client-key", "certs/client/client.key", "Path to client private key file (PEM format)")
	serverCACertPath := flag.String("server-ca", "certs/ca/ca.crt", "Path to server CA certificate file (PEM format)")
	flag.Parse()

	// Load client certificate and key
	clientCert, err := tls.LoadX509KeyPair(*clientCertPath, *clientKeyPath)
	if err != nil {
		log.Fatalf("Failed to load client certificate: %v", err)
	}

	// Load server's CA certificate
	caCert, err := ioutil.ReadFile(*serverCACertPath)
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a credentials object using the client certificate and CA
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCertPool,
	})

	// Set up gRPC connection with the credentials
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	message := &pb.Message{
		Name: "sigsegv1989",
	}
	response, err := c.SayHello(context.Background(), message)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	log.Printf("Response: %v", response)
}
