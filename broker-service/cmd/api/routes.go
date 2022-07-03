package main

/*
chi is a lightweight, idiomatic and
composable router for building Go HTTP services.
 It's especially good at helping you write large REST API
 services that are kept maintainable as your project grows
 and changes.
*/
import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	//These funcs accept a path and a handler to respond to the path
	mux.Use(middleware.Heartbeat("/ping"))
	//Handler for Broker Service defined in handlers.go
	mux.Post("/", app.Broker)
	return mux
}
