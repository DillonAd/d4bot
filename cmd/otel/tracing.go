package otel

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

func InitTracing(ctx context.Context, collectorEndpoint string) func() {
	if collectorEndpoint == "" {
		log.Println("no opentelemetry collector endpoint")
		return func() {}
	}

	client := otlptracegrpc.NewClient([]otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(collectorEndpoint),
		otlptracegrpc.WithInsecure(),
	}...)
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Printf("error creating trace exporter: %v", err)
		return func() {}
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("d4bot"),
		),
		resource.WithFromEnv(),
	)

	if err != nil {
		log.Printf("error creating tracing resource: %v", err)
		return func() {}
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))
	otel.SetTracerProvider(tracerProvider)

	return func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			log.Printf("error stopping tracer provider: %v", err)
		}
	}
}

func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	tracer := otel.Tracer("d4bot")
	ctx = baggage.ContextWithoutBaggage(ctx)
	return tracer.Start(ctx, name)
}

func SpanError(span trace.Span, err error) {
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
}
