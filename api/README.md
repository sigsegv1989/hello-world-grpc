# Generating Go Code from Protobuf for gRPC

This repository includes the generated Go code files for gRPC communication, specifically:

- `api/helloworld/hello.pb.go`
- `api/helloworld/hello_grpc.pb.go`

These files are generated from the `api/hello.proto` Protobuf file using the `protoc` tool along with the Go plugins. The generated code provides the necessary structures and interfaces for seamless communication between the gRPC client and server.

## How to Generate the Go Code

To generate the Go code files from the `api/hello.proto` Protobuf file, follow these steps:

1. Ensure that you have the Protocol Buffers compiler (`protoc`) installed on your system. If not, you can download it from the official GitHub repository: [protobuf/protobuf](https://github.com/protocolbuffers/protobuf/releases)

2. Install the Go plugins for Protocol Buffers by running the following command:
```
   go install google.golang.org/protobuf/cmd/protoc-gen-go
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
3. Run the protoc command from the main directory with the Go plugins to generate the code files. Use the following command:
```
   PATH="${PATH}:${HOME}/go/bin" protoc --go_out=. --go-grpc_out=. ./api/hello.proto
```

The generated hello.pb.go file contains the message definitions, and the hello_grpc.pb.go file contains the service definitions needed for gRPC communication.

*Keep in mind that if you make changes to the api/hello.proto file, you should re-run the protoc command to regenerate the code files to ensure that your changes are reflected in the generated Go code.*
