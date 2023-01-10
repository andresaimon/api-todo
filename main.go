package main

import (
	"fmt"
	"net/http"

	"github.com/andresaimon/api-todo/configs"
	"github.com/andresaimon/api-todo/handlers"
	"github.com/go-chi/chi/v5"
)

// Criação do server:
func main() {

	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	// declaração das rotas:
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{done}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
