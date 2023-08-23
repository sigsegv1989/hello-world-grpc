# Makefile for gRPC Server and Client

# Variables
OUTPUT_DIR:= bin
SERVER_BINARY := $(OUTPUT_DIR)/server
CLIENT_BINARY := $(OUTPUT_DIR)/client
SERVER_IMAGE := test/hello-world-grpc-server:latest
CLIENT_IMAGE := test/hello-world-grpc-client:latest

# Build Rules
build: build-server build-client

build-server:
	CGO_ENABLED=0 go build -o $(SERVER_BINARY) ./server

build-client:
	CGO_ENABLED=0 go build -o $(CLIENT_BINARY) ./client

# Docker Rules
docker-build: docker-build-server docker-build-client

docker-build-server: build-server
	docker build -t $(SERVER_IMAGE) -f server/Dockerfile.server .

docker-build-client: build-client
	docker build -t $(CLIENT_IMAGE) -f client/Dockerfile.client .

# Docker run rule
docker-run: docker-run-server docker-run-client

docker-run-server:
	docker run -p 50051:50051 --name server-container -d $(SERVER_IMAGE)

docker-run-client:
	docker run --network host -it --rm $(CLIENT_IMAGE)

# Clean Rule
clean:
	rm -rf $(OUTPUT_DIR)

# Phony Targets
.PHONY: build-server build-client docker-build-server docker-build-client docker-run-server docker-run-client clean
