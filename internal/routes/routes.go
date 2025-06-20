package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/joaquinbian/workout-api-go/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewMux()

	r.Get("/health", app.HealthCheck)

	return r
}
