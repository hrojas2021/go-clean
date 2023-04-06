package rest

import (
	"fmt"
	"net/http"

	"github.com/hugo.rojas/custom-api/internal/http/rest/handlers"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

// InitRoutes mounts all defaut routes
func InitRoutes(service iface.Service) *bunrouter.CompatRouter {
	// https://bunrouter.uptrace.dev/guide/golang-router.html#installation
	r := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
		bunrouter.Use(handlers.CorsMiddleware),
		bunrouter.Use(handlers.Telemetry),
		bunrouter.WithNotFoundHandler(notFoundHandler),
		bunrouter.WithMethodNotAllowedHandler(methodNotAllowedHandler),
	).Compat()

	resp := new(DefaultResp)
	h := handlers.New(service, resp)

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "API - Running OK")
	})

	r.POST("/login", h.Login)

	api := r.NewGroup("/api", bunrouter.Use(h.Authenticate))
	api.WithGroup("/", func(g *bunrouter.CompatGroup) {
		g.GET("/users", h.ListUsers)
		g.POST("/rooms", h.SaveRoom)
	})

	return r
}

func notFoundHandler(w http.ResponseWriter, req bunrouter.Request) error {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(
		w,
		"<html>BunRouter can't find a route that matches <strong>%s</strong></html>",
		req.URL.Path,
	)
	return nil
}

func methodNotAllowedHandler(w http.ResponseWriter, req bunrouter.Request) error {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(
		w,
		"<html>BunRouter does have a route that matches <strong>%s</strong>, "+
			"but it does not handle method <strong>%s</strong></html>",
		req.URL.Path, req.Method,
	)
	return nil
}
