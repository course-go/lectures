package main

import (
	"context"
	"net/http"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

// START RESOURCE OMIT

func newResource() (*resource.Resource, error) {
	return resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("My Awesome App"),
		),
	)
}

// END RESOURCE OMIT
// START TRACE-PROVIDER OMIT

func newTraceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	return otlptracehttp.New(ctx)
}

func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	r, _ := newResource()
	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}

// END TRACE-PROVIDER OMIT
// START METRIC-PROVIDER OMIT

func newMeterExporter(ctx context.Context) (sdkmetric.Exporter, error) {
	return otlpmetrichttp.New(ctx)
}

func newMeterProvider(exp sdkmetric.Exporter) *sdkmetric.MeterProvider {
	r, _ := newResource()
	return sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(r),
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exp)),
	)
}

// END METRIC-PROVIDER OMIT
// START TRACE-INIT OMIT

var (
	tracer trace.Tracer
	meter  metric.Meter
)

func main() {
	ctx := context.Background()

	te, _ := newTraceExporter(ctx)
	tp := newTraceProvider(te)
	defer tp.Shutdown(ctx)
	otel.SetTracerProvider(tp)
	tracer = tp.Tracer("My Awesome App")

	me, _ := newMeterExporter(ctx)
	mp := newMeterProvider(me)
	defer mp.Shutdown(ctx)
	otel.SetMeterProvider(mp)
	meter = mp.Meter("Ma Awesome App")
}

// END TRACE-INIT OMIT
// START CREATING-SPAN OMIT

func httpHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "my-awesome-span")
	defer span.End()
}

// END CREATING-SPAN OMIT
// START SPAN-ATTRIBUTES OMIT

func httpHandlerWithAttributes(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(context.Background(), "span-with-attributes", trace.WithAttributes(attribute.String("hello", "gophers")))
	span.SetAttributes(attribute.Bool("authorized", true), attribute.String("username", "bob"))
	defer span.End()
}

// END SPAN-ATTRIBUTES OMIT
// START SPAN-STATUS OMIT

func httpHandlerWithStatus(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "fluffy-gophers")
	defer span.End()
	gophers, err := getFluffyGophers(ctx)
	if err != nil {
		span.SetStatus(codes.Error, "getting fluffy gophers failed")
		span.RecordError(err)
	}
}

// END SPAN-STATUS OMIT

func getFluffyGophers(ctx context.Context) (int, error) {
	return 0, nil
}

// START METRIC OMIT

var apiCounter metric.Int64UpDownCounter

func init() {
	apiCounter, _ = meter.Int64UpDownCounter("api.request.counter",
		metric.WithDescription("HTTP request count."),
		metric.WithUnit("{call}"),
	)
	meter.Int64ObservableGauge("memory.heap",
		metric.WithDescription("Memory usage of the allocated heap objects."),
		metric.WithUnit("By"),
		metric.WithInt64Callback(func(_ context.Context, o metric.Int64Observer) error {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			o.Observe(int64(m.HeapAlloc))
			return nil
		}),
	)
}

func httpHandlerWithCounter(w http.ResponseWriter, r *http.Request) {
	apiCounter.Add(r.Context(), 1, metric.WithAttributes(semconv.HTTPStatusCode(http.StatusOK)))
}

// END METRIC OMIT
