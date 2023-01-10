package rest

import (
	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/http/rest/handlers"
	"github.com/hugo.rojas/custom-api/internal/http/rest/middlewares"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

// InitRoutes mounts all defaut routes
func InitRoutes(service iface.Service, conf *conf.Configuration) *bunrouter.CompatRouter {
	// https://bunrouter.uptrace.dev/guide/golang-router.html#installation
	r := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
		// add default error routes
	).Compat()

	h := handlers.New(service)

	r.POST("/login", h.Login)

	api := r.NewGroup("/api", bunrouter.Use(middlewares.Authenticate))
	api.WithGroup("/", func(g *bunrouter.CompatGroup) {
		g.GET("/users", h.ListUsers)
		g.POST("/rooms", h.SaveRoom)
	})

	return r
}
