package main

import (
	"fmt"

	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/http/rest"
	"github.com/hugo.rojas/custom-api/internal/infrastructure/api"
	"github.com/hugo.rojas/custom-api/internal/infrastructure/bootstrap"
	"github.com/hugo.rojas/custom-api/internal/infrastructure/db"
	"github.com/hugo.rojas/custom-api/internal/infrastructure/server"
)

func main() {
	config := conf.LoadViperConfig()
	db := db.InitDB(config)

	api := api.NewAPI(db, config)
	api.Handler = rest.InitRoutes(api)

	addr := fmt.Sprintf("%v:%v", "", config.PORT)

	b := bootstrap.NewBootstrap(api)
	bootstrap.InitServices(b)

	srv := server.NewServer(api.Handler, addr)
	srv.ListenAndServe()
}
