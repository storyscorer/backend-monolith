package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `toml:"server"`
}

type ServerConfig struct {
	Port int `toml:"port"`
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

type Handler struct {
	Router *mux.Router
}

func (h *Handler) Handle(port int) {
	h.Router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), h.Router); err != nil {
		fmt.Printf("Failed to open up the port %d\n", port)
	}
}

func main() {
	env := "local"
	if parsedEnv := os.Getenv("ENVIRONMENT"); parsedEnv != "" {
		env = parsedEnv
	}

	cfg := loadConfig(env)

	rtr := mux.NewRouter()
	h := Handler{
		Router: rtr,
	}

	h.Handle(cfg.Server.Port)
}
