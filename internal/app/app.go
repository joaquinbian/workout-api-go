package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joaquinbian/workout-api-go/internal/api"
	"github.com/joaquinbian/workout-api-go/internal/middleware"
	"github.com/joaquinbian/workout-api-go/internal/store"
	"github.com/joaquinbian/workout-api-go/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	UserHandler    *api.UserHandler
	TokenHandler   *api.TokenHandler
	Middleware     middleware.UserMiddleware
	DB             *sql.DB
}

func NewApplication() (*Application, error) {

	//vamos a usar logger para los logs pq nos da un mejor manejo de ellos y nos ayuda a saber que esta pasando
	//con logger se puede manejar mejor diferentes tipos de logs como errores, logs para debugging, etc
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	//store
	db, err := store.Open()
	if err != nil {
		return nil, err
	}
	err = store.MigrateFS(db, migrations.FS, ".")

	if err != nil {
		panic(err)
	}

	workoutStore := store.NewPostgresWorkoutStore(db)
	userStore := store.NewPostgresUserStore(db)
	tokenStore := store.NewPostgresTokenStore(db)

	//handlers
	workoutHandler := api.NewWorkoutHandler(workoutStore, logger)
	userHandler := api.NewUserHandler(userStore, logger)
	tokenHandler := api.NewTokenHander(tokenStore, userStore, logger)
	middlewareHandler := middleware.UserMiddleware{UserStore: userStore}

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		UserHandler:    userHandler,
		TokenHandler:   tokenHandler,
		Middleware:     middlewareHandler,
		DB:             db,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	//w: interface usada por HTTP handlers para crear respuestas HTTP
	//	con el contestamos al cliente

	//r: HTTP request recibida por el servidor, lo que nos envia el cliente

	fmt.Fprintln(w, "Server is up and running")
}
