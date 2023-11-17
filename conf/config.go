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

// JobsConfiguration holds all jobs tasks configuration
type JobsConfiguration struct {
	Enabled      bool
	Concurrency  int
	RedisAddress string
}

// LoggerConfiguration holds all logger configuration
type LoggerConfiguration struct {
	OUTPUT_PATH     string
	ERR_OUTPUT_PATH string
	DEFAULT_PATH    string
	MAX_SIZE        int
	MAX_BACKUPS     int
	MAX_AGE         int
}

// Configuration holds all configuration for this project
type Configuration struct {
	JWT            SecuriyConfiguration
	PORT           int `default:"9000"`
	DB             DBConfiguration
	LOGGER         LoggerConfiguration
	SERVER_TIMEOUT int
	Jobs           JobsConfiguration
	Telemetry      TelemetryConfiguration
	IS_PRODUCTION  bool
}

func LoadViperConfig() *Configuration {
	// viper.AutomaticEnv()
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("../conf")
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
