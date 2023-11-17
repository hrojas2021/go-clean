package conf

import (
	"log"

	"github.com/spf13/viper"
)

// DBConfiguration holds all the database related configuration.
type DBConfiguration struct {
	URL      string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Ssl      string
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
type SecurityConfiguration struct {
	Secret     string
	Expiration int
}

// JobsConfiguration holds all jobs tasks configuration
type JobsConfiguration struct {
	Enabled      bool
	Concurrency  int
	RedisAddress string
}

// LoggerConfiguration holds all logger configuration
type LoggerConfiguration struct {
	OutputPath    string
	ErrOutputPath string
	DefaultPath   string
	MaxSize       int
	MaxBackups    int
	MaxAge        int
}

// Configuration holds all configuration for this project
type Configuration struct {
	Jwt           SecurityConfiguration
	Port          int `default:"9000"`
	DB            DBConfiguration
	Logger        LoggerConfiguration
	ServerTimeout int
	Jobs          JobsConfiguration
	Telemetry     TelemetryConfiguration
	IsProduction  bool
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
