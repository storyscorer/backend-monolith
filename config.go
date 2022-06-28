package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	LogLevel string

	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string `toml:"dbname"`
}

func loadConfig(env string) (*Config, error) {
	viper.SetConfigFile(fmt.Sprintf("configs/%s.toml", env))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to load config file, %s", err.Error())
	}

	vip := viper.GetViper()
	var cfg Config
	err = vip.Unmarshal(&cfg)
	return &cfg, err
}
