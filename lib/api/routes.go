package api

import (
	"github.com/arcanjo96/golang-api-in-memory/lib/api/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(server *chi.Mux) {
	server.Route("/api", func(router chi.Router) {
		router.Get("/users", handlers.GetUsers)
		router.Get("/users/{id}", handlers.GetUser)
		router.Post("/users", handlers.CreateUser)
		router.Patch("/users/{id}", handlers.UpdateUser)
		router.Delete("/users/{id}", handlers.DeleteUser)
	})
}
