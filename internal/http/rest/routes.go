package rest

import (
	"fmt"
	"net/http"

	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/http/rest/handlers"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

// InitRoutes mounts all defaut routes
func InitRoutes(service iface.Service, conf *conf.Configuration) *bunrouter.CompatRouter {
	// https://bunrouter.uptrace.dev/guide/golang-router.html#installation
	r := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
		bunrouter.Use(handlers.CorsMiddleware),
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
