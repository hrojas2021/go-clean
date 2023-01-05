package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	srv *http.Server
}

func NewServer(r *httprouter.Router, h string) server {
	return server{
		srv: &http.Server{
			Addr:    h,
			Handler: r,
		},
	}
}

func (s *server) ListenAndServe() {

	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			log.Fatal("listen: ", err)
		}
	}()

	log.Printf("server listening on address %s\n", s.srv.Addr)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
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
