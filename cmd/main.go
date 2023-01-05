package main

import (
	"fmt"
	"log"

	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/infrastructure"
	"github.com/spf13/viper"
)

func loadViperConfig() *conf.Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app/config")
	viper.SetConfigType("yaml")
	var configuration conf.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &configuration
}

func main() {
	config := loadViperConfig()
	db := infrastructure.InitDB(config)

	api := infrastructure.NewAPI(db, config)
	api.Handler = infrastructure.InitRoutes(api)

	addr := fmt.Sprintf("%v:%v", "", config.PORT)

	b := infrastructure.NewBootstrap(api)
	infrastructure.InitServices(b)

	srv := infrastructure.NewServer(api.Handler, addr)
	srv.ListenAndServe()
}
