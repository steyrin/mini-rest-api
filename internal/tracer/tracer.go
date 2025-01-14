package tracer

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.17.0"
)

// InitTracer initializes the tracer provider with an exporter
func InitTracer(exportToConsole bool) func(context.Context) error {
	var exporter sdktrace.SpanExporter
	var err error

	if exportToConsole {
		exporter, err = stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			log.Fatalf("Failed to create console exporter: %v", err)
		}
	} else {
		exporter, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
		if err != nil {
			log.Fatalf("Failed to create Jaeger exporter: %v", err)
		}
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("Book Catalog API"),
		)),
	)

	otel.SetTracerProvider(tracerProvider)

	return tracerProvider.Shutdown
}
