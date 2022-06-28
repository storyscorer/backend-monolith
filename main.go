package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

func main() {
	env := "local"
	if parsedEnv := os.Getenv("ENVIRONMENT"); parsedEnv != "" {
		env = parsedEnv
	}

	cfg := loadConfig(env)

	rtr := mux.NewRouter()
	h := Handler{
		Config: cfg,
		Router: rtr,
	}

	h.Handle(cfg.Server.Port)
}
