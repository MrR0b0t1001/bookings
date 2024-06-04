package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/MrR0b0t1001/bookings/pkg/config"
	"github.com/MrR0b0t1001/bookings/pkg/handlers"
)

func routing(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter() // using chi module

	mux.Use(
		middleware.Recoverer,
	) // Recovers and Logs a Panic event and returns an HTTP 500 (Internal Server Error)

	mux.Use(NoSurf)      // We create a CSRF token -> Cross-Site-Request-Forgery
	mux.Use(SessionLoad) // We Load and Save the state of the current session

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	log.Println("\nProduction Mode: ", app.InProduction)

	return mux
}
