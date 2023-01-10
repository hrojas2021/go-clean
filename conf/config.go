package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
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

// Configuration holds all openTelemetry configuration
type TelemetryConfiguration struct {
	Enabled   bool
	Name      string
	JaegerURL string
	Version   string
	FilePath  string
}

// Configuration holds all jwt configuration
type SecuriyConfiguration struct {
	SECRET string
}

// Configuration holds all configuration for this project
type Configuration struct {
	JWT       SecuriyConfiguration
	PORT      int `default:"9000"`
	DB        DBConfiguration
	Telemetry TelemetryConfiguration
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

func LoadViperConfig() *Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("/app/config")
	viper.SetConfigType("yaml")
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &configuration
}
