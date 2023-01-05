package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/hugo.rojas/custom-api/internal/iface"
	"github.com/julienschmidt/httprouter"
)

func panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println(r.URL.Path, string(debug.Stack()))
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
}

// InitRoutes mounts all defaut routes
func InitRoutes(service iface.Service) *httprouter.Router {
	r := httprouter.New()
	/********************** GLOBAL OPTIONS *****************/
	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}
		w.WriteHeader(http.StatusNoContent)
	})

	/********************** PUBLIC ROUTES *****************/
	r.GET("/healthcheck", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("API running - OK")
	})

	/********************** DEFAULT ERROR  ROUTES *****************/
	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	})

	r.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	r.PanicHandler = panicHandler

	/********************** GROUP ROUTES *****************/
	return r
}
