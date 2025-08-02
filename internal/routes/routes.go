package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/joaquinbian/workout-api-go/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.GetWorkoutByID)
	r.Get("/workouts", app.WorkoutHandler.GetWorkouts)

	r.Post("/workouts", app.WorkoutHandler.CreateWorkout)
	return r
}
