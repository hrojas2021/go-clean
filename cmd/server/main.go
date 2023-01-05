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
	"github.com/hugo.rojas/custom-api/internal/infrastructure/db"
	"github.com/hugo.rojas/custom-api/internal/io"
	"github.com/hugo.rojas/custom-api/internal/io/database"
	"github.com/hugo.rojas/custom-api/internal/service"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	srv *http.Server
}

func newServer(r *httprouter.Router, h string) server {
	return server{
		srv: &http.Server{
			Addr:    h,
			Handler: r,
		},
	}
}

func main() {
	config := conf.LoadViperConfig()
	db := db.InitDB(config)

	io := io.New(database.New(db))
	service := service.New(config, io)

	r := rest.InitRoutes(service)

	addr := fmt.Sprintf("%v:%v", "", config.PORT)

	srv := newServer(r, addr)
	listenAndServe(&srv)
}

func listenAndServe(s *server) {

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
