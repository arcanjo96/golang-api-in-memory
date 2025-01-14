package main

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/arcanjo96/golang-api-in-memory/lib/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	mux := chi.NewMux()
	server := http.Server{
		Addr:                         ":8000",
		Handler:                      mux,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    &tls.Config{},
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  1 * time.Minute,
	}

	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	api.RegisterRoutes(mux)

	print("Server is running on port 8000")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
