# Certificate and Key Generation for gRPC Communication

This guide outlines the steps to generate certificates and keys required for enabling secure communication and authentication between the gRPC client and server.

## Steps to Generate Certificate and Key

1. **Generate CA Certificate and Key:**
   - Run the following commands to generate the CA certificate and key:
     ```
     openssl genpkey -algorithm RSA -out ca/ca.key
     openssl req -x509 -new -key ca/ca.key -out ca/ca.crt -subj "/C=US/ST=California/L=San Francisco/O=My Organization/OU=IT Department/CN=My CA/emailAddress=ca@example.com"
     ```

2. **Generate Server Certificate and Key:**
   - Run the following commands to generate the server certificate and key:
     ```
     openssl genpkey -algorithm RSA -out server/server.key
     openssl req -new -key server/server.key -out server/.server.csr -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost"
     openssl x509 -req -extfile <(printf "subjectAltName=DNS:localhost") -in server/.server.csr -CA ca/ca.crt -CAkey ca/ca.key -CAcreateserial -out server/server.crt
     ```

3. **Generate Client Certificate and Key:**
   - Run the following commands to generate the client certificate and key:
     ```
     openssl genpkey -algorithm RSA -out client/client.key
     openssl req -new -key client/client.key -out client/.client.csr -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost"
     openssl x509 -req -extfile <(printf "subjectAltName=DNS:localhost") -in client/.client.csr -CA ca/ca.crt -CAkey ca/ca.key -CAcreateserial -out client.crt
     ```

## Generated Files Description

- `ca/ca.crt`: CA certificate used to sign server and client certificates.
- `ca/ca.key`: Private key of the CA certificate.
- `server/server.crt`: Server's certificate signed by the CA.
- `server/server.key`: Private key of the server's certificate.
- `client/client.crt`: Client's certificate signed by the CA.
- `client/client.key`: Private key of the client's certificate.

**Note**: Keep the generated certificates and keys secure and do not embed them in Docker images. It's recommended to handle sensitive data separately and securely.

These certificates and keys are essential for establishing a secure and authenticated connection between the gRPC client and server.
