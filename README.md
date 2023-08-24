# Hello World gRPC

This repository contains a simple example of a client-server application using gRPC in Go (golang). The application demonstrates how to exchange personalized greetings and messages between a client and a server.

## Prerequisites

Before running the code in this repository, make sure you have the following prerequisites installed:

- Go (golang): You can download and install Go from the official website: https://golang.org/dl/
- Protocol Buffers (Protobuf): You can download and install from the official GitHub repository: [protobuf/protobuf](https://github.com/protocolbuffers/protobuf/releases)
- gRPC Go packages: You can install them using the following command:
```
   go install google.golang.org/protobuf/cmd/protoc-gen-go
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## How to Run

1. Clone this repository to your local machine:
```
   git clone https://github.com/sigsegv1989/hello-world-grpc.git
```

2. Navigate to the repository directory:
```
   cd hello-world-grpc
```

3. Build the client, server binary and docker images:
```
   make docker-build
```

4. Start the server:
```
   make docker-run-server
```

5. Run the client:
```
   make docker-run-client
```

6. You should see the interaction between the client and server, exchanging personalized greetings and messages.

## Project Structure

- `api/hello.proto`: Protocol Buffers (.proto) file defining the message and service.
- `server/server.go`: Server implementation that receives and responds to gRPC requests.
- `client/client.go`: Client implementation that sends gRPC requests to the server.
- `certs/`: TLS/SSL certificates to enable the mutual authentication between gRPC client and server.

## License

This repository contains a simple client-server application intended for study purposes. Please note that this code is provided as-is and comes with no guarantee of functioning correctly in a production environment.

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
