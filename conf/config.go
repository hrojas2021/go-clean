package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// DBConfiguration holds all the database related configuration.
type DBConfiguration struct {
	URL      string
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	NAME     string
	SSL      string
}

// Configuration holds all configuration for this project
type Configuration struct {
	PORT int `default:"9500"`
	DB   DBConfiguration
}

func loadEnvironment(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Load(filename)
	} else {
		err = godotenv.Load()

		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}

func LoadConfig(filename string) *Configuration {
	if err := loadEnvironment(filename); err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	config := new(Configuration)
	if err := envconfig.Process("BF", config); err != nil {
		log.Fatalf("Failed to process configuration: %s", err)
	}

	return config
}
