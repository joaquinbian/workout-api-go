package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application, error) {

	//vamos a usar logger para los logs pq nos da un mejor manejo de ellos y nos ayuda a saber que esta pasando
	//con logger se puede manejar mejor diferentes tipos de logs como errores, logs para debugging, etc
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	app := &Application{
		Logger: logger,
	}

	return app, nil
}
