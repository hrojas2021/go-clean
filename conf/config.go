package conf

import (
	"log"

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
	SECRET     string
	EXPIRATION int
}

// Configuration holds all configuration for this project
type Configuration struct {
	JWT       SecuriyConfiguration
	PORT      int `default:"9000"`
	DB        DBConfiguration
	Telemetry TelemetryConfiguration
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
