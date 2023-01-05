package main

import (
	"context"
	"os"

	"github.com/hugo.rojas/custom-api/conf"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// Returns a new instance of TracerProvider for OpenTelemetry use
func newTraceProvider(ctx context.Context, cfgTelemetry conf.TelemetryConfiguration) (*trace.TracerProvider, error) {
	var err error
	var exporter trace.SpanExporter
	if cfgTelemetry.JaegerURL != "" {
		exporter, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfgTelemetry.JaegerURL)))
		if err != nil {
			return nil, err
		}
	} else {
		w := os.Stderr
		if cfgTelemetry.FilePath != "" {
			w, err = os.OpenFile(cfgTelemetry.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
			if err != nil {
				return nil, err
			}
			go func() {
				<-ctx.Done()
				_ = w.Close()
			}()
		}

		exporter, err = stdouttrace.New(
			stdouttrace.WithWriter(w),
			stdouttrace.WithPrettyPrint(),
			stdouttrace.WithoutTimestamps(),
		)
		if err != nil {
			return nil, err
		}
	}

	re := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(cfgTelemetry.Name),
		attribute.String("version", cfgTelemetry.Version),
	)

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exporter),
		trace.WithResource(re),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}))

	return tp, nil
}
