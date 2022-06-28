package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
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

func loadConfig(env string) *Config {
	viper.SetConfigFile(fmt.Sprintf("configs/%s.toml", env))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to load config file, %s", err.Error())
	}

	vip := viper.GetViper()
	var cfg Config
	err = vip.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Failed to unmarshal the config file, %s", err.Error())
	}

	return &cfg
}
