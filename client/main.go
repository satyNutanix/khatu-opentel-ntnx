package main

import (
  "context"
  "log"

  "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
  "go.opentelemetry.io/otel"
  "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
  "go.opentelemetry.io/otel/sdk/resource"
  "go.opentelemetry.io/otel/sdk/trace"

  pb "shyam-opentel/example"

  "google.golang.org/grpc"
)

func main() {
  // Setup tracing
  tracerProvider := setupTracing()
  defer tracerProvider.Shutdown(context.Background())

  // Connect to the server
  conn, err := grpc.Dial(
    "localhost:50051",
    grpc.WithInsecure(),
    grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
  )
  if err != nil {
    log.Fatalf("Failed to connect: %v", err)
  }
  defer conn.Close()

  client := pb.NewExampleServiceClient(conn)

  // Create a context with a span for tracing
  tracer := otel.Tracer("client")
  ctx, span := tracer.Start(context.Background(), "SayHelloRequest")
  defer span.End()

  // Log the trace ID from the client
  log.Printf("Client TraceID: %s", span.SpanContext().TraceID())

  // Send the gRPC request
  resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "OpenTelemetry"})
  if err != nil {
    log.Fatalf("Error calling SayHello: %v", err)
  }

  log.Printf("Response from server: %s", resp.Message)
}

func setupTracing() *trace.TracerProvider {
  exporter, _ := stdouttrace.New(stdouttrace.WithPrettyPrint())
  return trace.NewTracerProvider(
    trace.WithBatcher(exporter),
    trace.WithResource(resource.NewSchemaless()),
  )
}
