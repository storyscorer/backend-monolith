package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Handler struct {
	Config *Config
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

func createLogger(logLvl string) (*zap.Logger, error) {
	lvl := zap.NewAtomicLevel()

	err := lvl.UnmarshalText([]byte(logLvl))
	if err != nil {
		lvl.SetLevel(zapcore.InfoLevel)
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl
	lgr, err := cfg.Build()
	if err != nil {
		log.Fatal("Failed to instansiate the logger")
	}

	return lgr, nil
}

func main() {
	env := "local"
	if parsedEnv := os.Getenv("ENVIRONMENT"); parsedEnv != "" {
		env = parsedEnv
	}

	cfg, err := loadConfig(env)
	if err != nil {
		log.Fatalf("Failed to load the config, terminating: %s\n", err.Error())
	}

	lgr, err := createLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Failed to instansiate the logger, terminating: %s\n", err.Error())
	}
	unbindLgr := zap.ReplaceGlobals(lgr)
	defer unbindLgr()

	rtr := mux.NewRouter()
	h := Handler{
		Config: cfg,
		Router: rtr,
	}

	h.Handle(cfg.Server.Port)
}
