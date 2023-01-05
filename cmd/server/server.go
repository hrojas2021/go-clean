package main

import (
	"net/http"

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
