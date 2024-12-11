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

# Use a base image for the final container
FROM golang:1.22

# Install any dependencies required for your app (if needed)
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Create the directory structure for the config file
RUN mkdir -p /home/nutanix/config/observability

# Copy the config.json file to the desired location
COPY config.json /home/nutanix/config/observability/config.json

# Copy server and client binaries from the builder stage
COPY --from=builder /app/server /server
COPY --from=builder /app/client /client

# Debug step to list files in the final image
RUN ls -l /home/nutanix/config/observability

# Set the working directory to where the config file resides
WORKDIR /home/nutanix/config/observability

# Set the entrypoint to the server binary
ENTRYPOINT ["/server/main"]