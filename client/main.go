package main

import (
	"context"
	"log"

	pb "shyam-opentel/example"

	"google.golang.org/grpc"

	"github.com/nutanix-core/go-cache/util-go/tracer"
)

func main() {
	closer := tracer.InitTracer("Foundation-Central")
	if closer != nil {
		// The tracer needs to be closed before the service stops, this will clear the traces stored in the cache.
		defer closer.Close()
	}

	unaryTrace, streamTrace := tracer.GrpcRequestTraceOptions()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	if unaryTrace != nil {
		opts = append(opts, grpc.WithUnaryInterceptor(unaryTrace))
	}
	if streamTrace != nil {
		opts = append(opts, grpc.WithStreamInterceptor(streamTrace))
	}

	// Connect to the server
	conn, err := grpc.Dial(
		"localhost:50051",
		opts...,
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewExampleServiceClient(conn)

  spn, ctx:= tracer.StartSpan(context.Background(), "callEampleMethod")
  log.Printf("Root Client TraceID: %s", spn)

	span, ctx := tracer.StartRpcClientSpan(nil, "callEampleMethod", ctx)
	log.Printf("Client TraceID: %s", span)
	defer span.Finish()

	// Send the gRPC request with the context
	resp, err := client.SayHello(ctx, &pb.HelloRequest{
		Name:    "OpenTelemetry",
	})

	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}

	log.Printf("Response from server: %s", resp.Message)

}
