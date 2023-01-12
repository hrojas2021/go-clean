package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

/* type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
} */

func main() {
	http.Handle("/other", http.HandlerFunc(handleIndex))
	http.HandleFunc("/api", handleIndex)

	r := httprouter.New()
	r.HandleMethodNotAllowed = true
	r.GET("/login", indexHandler) // path string, handle httprouter.Handle)

	/*
		r.Handle()                    // method string, path string, handle httprouter.Handle)
		r.Handler()                   // method string, path string, handler http.HandlerFunc)
		r.HandlerFunc()               // method string, path string, handler http.HandlerFunc)
	*/

	logger, _ := zap.NewProduction()
	logger.Info("INFO log level message")
	logger.Warn("Warn log level message")
	logger.Error("Error log level message")

	sug := logger.Sugar()
	_ = sug

}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func handleIndex(w http.ResponseWriter, r *http.Request) {

}
