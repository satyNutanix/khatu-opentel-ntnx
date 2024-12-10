package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nutanix-core/go-cache/util-go/tracer"
	"google.golang.org/grpc"

	pb "shyam-opentel/example"
)

type server struct {
	pb.UnimplementedExampleServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	span, ctx := tracer.StartRpcServerSpan(nil, "SayHello", ctx)
	//span.AddEvent("Processing SayHello")
	span.SetTag("exampleTag2", "exampleValue2")
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}

func main() {
	closer := tracer.InitTracer("Foundation-Central")
	if closer != nil {
		// The tracer needs to be closed before the service stops, this will clear the traces stored in the cache.
		defer closer.Close()
	}


  unaryTrace, streamTrace := tracer.GrpcServerTraceOptions()
  var opts []grpc.ServerOption
  if unaryTrace != nil {
      unaryOpt := grpc.UnaryInterceptor(unaryTrace)
      opts = append(opts, unaryOpt)
  }
  if streamTrace != nil {
      streamOpt := grpc.StreamInterceptor(streamTrace)
      opts = append(opts, streamOpt)
  }
  grpcServer := grpc.NewServer(opts...)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pb.RegisterExampleServiceServer(grpcServer, &server{})

	log.Println("gRPC Server running on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
