package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/http/rest"
	"github.com/hugo.rojas/custom-api/internal/io"
	"github.com/hugo.rojas/custom-api/internal/io/database"
	"github.com/hugo.rojas/custom-api/internal/service"
	"github.com/hugo.rojas/custom-api/internal/telemetry"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {

	ctx, shutdown := context.WithCancel(context.Background())
	defer shutdown()
	config := conf.LoadViperConfig()
	db := database.InitDB(config)

	io := io.New(database.New(db))

	var tp *trace.TracerProvider
	if config.Telemetry.Enabled {
		tp, err := newTraceProvider(ctx, config.Telemetry)
		if err != nil {
			log.Fatal("Error converting the steps: ", err.Error())
		}
		io = telemetry.NewIO(io, tp)
	}

	service := service.New(config, io)
	if config.Telemetry.Enabled {
		service = telemetry.NewService(service, tp)
	}

	r := rest.InitRoutes(service)
	addr := fmt.Sprintf("%v:%v", "", config.PORT)
	srv := newServer(r, addr)
	listenAndServe(ctx, &srv)
}

func listenAndServe(ctx context.Context, s *server) {

	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			log.Fatal("listen: ", err)
		}
	}()

	log.Printf("server listening on address %s\n", s.srv.Addr)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
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
