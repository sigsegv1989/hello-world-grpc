package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	pb "github.com/sigsegv1989/hello-world-grpc/api/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	// Define command-line flags with default values
	clientCertPath := flag.String("client-cert", "certs/client/client.crt", "Path to client certificate file (PEM format)")
	clientKeyPath := flag.String("client-key", "certs/client/client.key", "Path to client private key file (PEM format)")
	serverCACertPath := flag.String("server-ca", "certs/ca/ca.crt", "Path to server CA certificate file (PEM format)")
	targetFlag := flag.String("target", "localhost:50051", "Address of the gRPC server")
	namesFlag := flag.String("names", "SIGINT,SIGKILL,SIGQUIT,SIGSEGV,SIGTERM", "Comma-separated list of names")

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
	conn, err := grpc.Dial(*targetFlag, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	names := strings.Split(*namesFlag, ",")

	// Create a batch of Message objects
	var messages []*pb.Message
	for _, name := range names {
		messages = append(messages, &pb.Message{
			Name: name,
		})
	}

	request := &pb.BatchRequest{
		Requests: messages,
	}

	// The exponential backoff algorithm helps handle transient failures gracefully.
	// It combines the benefits of randomized intervals (jitter) and exponential
	// increase in intervals to avoid overloading the system during retries.
	// The algorithm calculates the maximum time spent on retries while considering
	// jitter and the exponential backoff sequence, ensuring resilience and efficiency.
	// If all retries fail within the maximum elapsed time, the operation exits with an error.

	// Set up exponential backoff algorithm for retries
	backOff := backoff.NewExponentialBackOff()
	backOff.InitialInterval = 1 * time.Second

	// Configure jitter (randomized delay) to avoid synchronization
	backOff.RandomizationFactor = 0.5

	// Configure backoff multiplier for exponential increase in intervals
	backOff.Multiplier = 2.0

	// Retries until maximum elapsed time
	backOff.MaxElapsedTime = 120 * time.Second

	var response *pb.BatchResponse
	operation := func() error {
		var err error
		// Send the request and handle the response
		response, err = c.SayHello(context.Background(), request)
		if err != nil {
			// Check if the error has a status code
			if statusErr, ok := status.FromError(err); ok {
				switch statusErr.Code() {
				case codes.InvalidArgument:
					// No retry on invalid argument.
					log.Fatalf("Invalid argument for the request, not going to retry: %v", err)
				case codes.Unavailable:
					// Retry in case of server unavailable
					return statusErr.Err()
				}
			}
			return err
		}

		log.Print("Server responses:")
		for _, resp := range response.Responses {
			log.Printf("Name: %s, Greeting: %s, Count: %d\n", resp.Name, resp.Greeting, resp.Count)
		}

		return nil
	}

	// Retry with exponential backoff
	if err := backoff.Retry(operation, backOff); err != nil {
		log.Fatalf("Max retries reached. Exiting... Last error: %v", err)
	}
}
