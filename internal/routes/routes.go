package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/joaquinbian/workout-api-go/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)
		r.Get("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.GetWorkoutByID))
		r.Get("/workouts", app.Middleware.RequireUser(app.WorkoutHandler.GetWorkouts))
		r.Post("/workouts", app.Middleware.RequireUser(app.WorkoutHandler.CreateWorkout))
		r.Put("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.UpdateWorkout))
		r.Delete("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.DeleteWorkout))
	})
	//WORKOUTS
	r.Get("/health", app.HealthCheck)
	r.Post("/users", app.UserHandler.HandleRegisterUser)
	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)
	return r
}
