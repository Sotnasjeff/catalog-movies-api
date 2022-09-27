package config

import (
	"context"
	"log"

	"github.com/spf13/viper"
)

var configuration *config

type config struct {
	DatabaseConfig DatabaseConfiguration
	ApiConfig      ApiConfiguration
}

type DatabaseConfiguration struct {
	Host     string
	Username string
	Password string
	Database string
	Port     string
}

type ApiConfiguration struct {
	Port string
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(context.Background(), err, "Fail to init configurations")
	}

	configuration = new(config)
	configuration.ApiConfig = ApiConfiguration{
		Port: viper.GetString("api.port"),
	}
	configuration.DatabaseConfig = DatabaseConfiguration{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.name"),
	}
}

func Config() *config {
	return configuration
}
