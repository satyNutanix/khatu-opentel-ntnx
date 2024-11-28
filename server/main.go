package main

import (
  "context"
  "fmt"
  "log"
  "net"

  "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
  "go.opentelemetry.io/otel"
  "go.opentelemetry.io/otel/propagation"
  "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
  "go.opentelemetry.io/otel/sdk/resource"
  sdktrace "go.opentelemetry.io/otel/sdk/trace"
  "go.opentelemetry.io/otel/trace"

  "google.golang.org/grpc"

  pb "shyam-opentel/example"
)

type server struct {
  pb.UnimplementedExampleServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
  span := trace.SpanFromContext(ctx)
  span.AddEvent("Received request in SayHello")
  log.Printf("TraceID: %s", span.SpanContext().TraceID())
  return &pb.HelloResponse{Message: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}

func main() {
  // Setup tracing
  tracerProvider := setupTracing()
  otel.SetTracerProvider(tracerProvider)
  otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
  defer tracerProvider.Shutdown(context.Background())

  listener, err := net.Listen("tcp", ":50051")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  grpcServer := grpc.NewServer(
    grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
  )

  pb.RegisterExampleServiceServer(grpcServer, &server{})

  log.Println("gRPC Server running on :50051")
  if err := grpcServer.Serve(listener); err != nil {
    log.Fatalf("Failed to serve: %v", err)
  }
}

func setupTracing() *sdktrace.TracerProvider {
  exporter, _ := stdouttrace.New(stdouttrace.WithPrettyPrint())
  return sdktrace.NewTracerProvider(
    sdktrace.WithBatcher(exporter),
    sdktrace.WithResource(resource.NewSchemaless()),
  )
}
