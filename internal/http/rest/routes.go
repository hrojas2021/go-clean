package rest

import (
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
		// add default error routes
	).Compat()

	// r.Use(Telemetry())

	h := handlers.New(service)
	// 	// r.Use(Recoverer(resp))  CHECK Recoverer and Check if Resp is needed

	r.WithGroup("/api/", func(g *bunrouter.CompatGroup) {
		g.GET("/campaigns/:id", h.GetCampaign)
	})

	return r
}
