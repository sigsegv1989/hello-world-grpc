# Hello World gRPC

This repository contains a simple example of a client-server application using gRPC in Go (golang). The application demonstrates how to exchange personalized greetings and messages between a client and a server.

## Prerequisites

Before running the code in this repository, make sure you have the following prerequisites installed:

- Go (golang): You can download and install Go from the official website: https://golang.org/dl/
- Protocol Buffers (Protobuf) and gRPC Go packages: You can install them using the following command:
```
go get google.golang.org/protobuf/cmd/protoc-gen-go \
       google.golang.org/grpc/cmd/protoc-gen-go-grpc
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

3. Start the server:
```
go run server/server.go
```

4. In a separate terminal window, run the client:
```
go run client/client.go
```

5. You should see the interaction between the client and server, exchanging personalized greetings and messages.

## Project Structure

- `api/hello.proto`: Protocol Buffers (.proto) file defining the message and service.
- `server/server.go`: Server implementation that receives and responds to gRPC requests.
- `client/client.go`: Client implementation that sends gRPC requests to the server.

## License

This repository contains a simple client-server application intended for study purposes. Please note that this code is provided as-is and comes with no guarantee of functioning correctly in a production environment.

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
