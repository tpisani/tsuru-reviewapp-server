package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type HealthcheckHandler struct{}

func (h HealthcheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "online",
	})
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Method(http.MethodGet, "/healthcheck/", HealthcheckHandler{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Running at :%s...\n", port)
	http.ListenAndServe(":"+port, r)
}
