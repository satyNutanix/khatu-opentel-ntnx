module shyam-opentel

go 1.22.7

toolchain go1.22.9

require (
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.57.0
	go.opentelemetry.io/otel v1.32.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.32.0
	go.opentelemetry.io/otel/sdk v1.32.0
	go.opentelemetry.io/otel/trace v1.32.0
	google.golang.org/grpc v1.68.0
	google.golang.org/protobuf v1.35.2
)

require (
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/glog v1.2.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/nutanix-core/go-cache v0.0.0-20241209093952-eb03f692b1e7 // indirect
	go.opentelemetry.io/contrib/propagators/jaeger v0.0.0-00010101000000-000000000000 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v0.0.0-00010101000000-000000000000 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241104194629-dd2ea8efbc28 // indirect
)

replace (
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc => go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.32.0
	go.opentelemetry.io/contrib/propagators/jaeger => go.opentelemetry.io/contrib/propagators/jaeger v1.7.0
	go.opentelemetry.io/otel => go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/jaeger => go.opentelemetry.io/otel/exporters/jaeger v1.7.0
	go.opentelemetry.io/otel/metric => go.opentelemetry.io/otel/metric v0.30.0
	go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v1.7.0
)
