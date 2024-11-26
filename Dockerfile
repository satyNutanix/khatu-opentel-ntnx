# Use the official Golang image to build the application
FROM golang:1.22 as builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . . 

# Build the server and client binaries
RUN go build -o server ./server/main.go
RUN go build -o client ./client/main.go

# Ensure server has execute permission (in the builder stage)
RUN chmod +x /app/server

# Debug step to list files in /app
RUN ls -l /app

# Use a base image that includes a shell for debugging
FROM debian:latest as debug

# Copy server and client from the builder stage
COPY --from=builder /app/server /server
COPY --from=builder /app/client /client

# Debug step to list files in the final image
RUN ls -l /

# Set the entrypoint to the server binary
ENTRYPOINT ["/server"]

# Use the minimal base image for the final image
FROM ubuntu:22.04

# Copy server and client from the builder stage
COPY --from=builder /app/server /server
COPY --from=builder /app/client /client

# Set the entrypoint to the server binary
ENTRYPOINT ["/server/main"]