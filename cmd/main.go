package main

import (
	"fmt"

	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/infrastructure"
)

func main() {
	var configFile = ""
	config := conf.LoadConfig(configFile)
	db := infrastructure.InitDB(config)

	api := infrastructure.NewAPI(db, config)
	api.Handler = infrastructure.InitRoutes(api)

	addr := fmt.Sprintf("%v:%v", "", config.PORT)

	b := infrastructure.NewBootstrap(api)
	infrastructure.InitServices(b)

	srv := infrastructure.NewServer(api.Handler, addr)
	srv.ListenAndServe()
}
