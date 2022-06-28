package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	LogLevel string `toml:"log_level"`

	Server   ServerConfig   `toml:"server"`
	Database DatabaseConfig `toml:"database"`
}

type ServerConfig struct {
	Port int `toml:"port"`
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
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
