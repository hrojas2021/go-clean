package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
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
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func handleIndex(w http.ResponseWriter, r *http.Request) {

}
