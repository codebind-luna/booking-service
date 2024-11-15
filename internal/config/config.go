package config

import (
	"fmt"
	"os"

	"github.com/codebind-luna/booking-service/internal/constants"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"server"`
	Log struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	} `mapstructure:"log"`
	Repository struct {
		Type string `mapstructure:"type"`
	} `mapstructure:"repository"`
	Inmemory struct {
		// no of seats each section (i.e A, B) will have
		Seats int `mapstructure:"seats"`
	}
}

func NewConfig() (*Config, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("../..")
	viper.SetConfigType("yaml")

	// Read in the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	viper.AutomaticEnv()

	viper.BindEnv("server.host", constants.EnvServerHost)
	viper.BindEnv("server.port", constants.EnvServerPort)
	viper.BindEnv("log.level", constants.EnvLogLevel)
	viper.BindEnv("log.format", constants.EnvLogFormat)
	viper.BindEnv("repository.type", constants.EnvRepositoryType)
	viper.BindEnv("inmemory.seats", constants.EnvInmemorySeats)

	// Set default values
	viper.SetDefault("server.host", constants.DefaultGRPCHost)
	viper.SetDefault("server.port", constants.DefaultGRPCPort)
	viper.SetDefault("log.level", constants.DefaultLogLevel)
	viper.SetDefault("log.format", constants.DefaultLogFormat)
	viper.SetDefault("repository.type", constants.DefaultRepositoryType)
	viper.SetDefault("inmemory.seats", constants.DefaultSeats)

	seats := os.Getenv(constants.EnvInmemorySeats)
	if seats == "" {
		os.Setenv(constants.EnvInmemorySeats, viper.GetString("inmemory.seats"))
	}

	logLevel := os.Getenv(constants.EnvLogLevel)
	if logLevel == "" {
		os.Setenv(constants.EnvLogLevel, viper.GetString("log.level"))
	}

	logFormat := os.Getenv(constants.EnvLogFormat)
	if logFormat == "" {
		os.Setenv(constants.EnvLogFormat, viper.GetString("log.format"))
	}

	port := viper.GetInt("server.port")

	if port < constants.MinPort || port > constants.MaxPort {
		return nil, fmt.Errorf("invalid port value: %d port must be between %d and %d", port, constants.MinPort, constants.MaxPort)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return &config, nil
}
