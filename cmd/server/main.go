package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/http/rest"
	"github.com/hugo.rojas/custom-api/internal/io"
	"github.com/hugo.rojas/custom-api/internal/io/database"
	"github.com/hugo.rojas/custom-api/internal/jobs"
	"github.com/hugo.rojas/custom-api/internal/logger"
	"github.com/hugo.rojas/custom-api/internal/service"
	"github.com/hugo.rojas/custom-api/internal/telemetry"
	"github.com/uptrace/bunrouter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
)

const (
	asyncJobs = "jobs"
)

type server struct {
	srv *http.Server
}

func newServer(r *bunrouter.CompatRouter, h string, timeout int) server {
	return server{
		srv: &http.Server{
			Addr:              h,
			Handler:           r,
			ReadHeaderTimeout: time.Duration(timeout) * time.Second,
		},
	}
}

func main() {
	ctx, shutdown := context.WithCancel(context.Background())
	defer shutdown()
	config := conf.LoadViperConfig()
	loggerConfig := zap.NewDevelopmentConfig()

	if config.IsProduction {
		loggerConfig = zap.NewProductionConfig()
		loggerConfig.OutputPaths = []string{config.Logger.OutputPath}
		loggerConfig.ErrorOutputPaths = []string{config.Logger.ErrOutputPath}
	}

	logger := logger.NewZapLogger(&config.Logger, loggerConfig, !config.IsProduction)
	defer func() {
		if err := logger.Sync(); err != nil {
			panic(err)
		}
	}()

	db := database.InitDB(config)

	io := io.New(database.New(db))
	var tp *trace.TracerProvider
	if config.Telemetry.Enabled {
		tp, err := newTraceProvider(ctx, config.Telemetry)
		if err != nil {
			logger.Fatal("Error converting the steps: " + err.Error())
		}
		io = telemetry.NewIO(io, tp)
	}

	service := service.New(config, io)
	if config.Telemetry.Enabled {
		service = telemetry.NewService(service, tp)
	}

	jobs.InitJobsQueue(*config, logger.Named(asyncJobs), service)

	r := rest.InitRoutes(service)
	addr := fmt.Sprintf("%v:%v", "", config.Port)
	srv := newServer(r, addr, config.ServerTimeout)
	listenAndServe(ctx, &srv)
}

func listenAndServe(ctx context.Context, s *server) {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			log.Fatal("listen: ", err)
		}
	}()

	log.Printf("server listening on address %s\n", s.srv.Addr)

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

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

	re, _ := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfgTelemetry.Name),
			attribute.String("version", cfgTelemetry.Version),
		),
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
