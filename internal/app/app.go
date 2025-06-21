package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joaquinbian/workout-api-go/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {

	//vamos a usar logger para los logs pq nos da un mejor manejo de ellos y nos ayuda a saber que esta pasando
	//con logger se puede manejar mejor diferentes tipos de logs como errores, logs para debugging, etc
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	//store

	//handlers
	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	//w: interface usada por HTTP handlers para crear respuestas HTTP
	//	con el contestamos al cliente

	//r: HTTP request recibida por el servidor, lo que nos envia el cliente

	fmt.Fprintln(w, "Server is up and running")
}
