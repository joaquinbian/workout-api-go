package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/joaquinbian/workout-api-go/internal/app"
)

func main() {

	app, err := app.NewApplication()

	if err != nil {
		//log.Fatal("An error ocurred instanciating the app")
		//la funcion panic va a crashear la app. Lo usamos para situaciones que NO DEBERIAN OCURRIR
		//si algo de esto ocurre, no hay nada que hacer, la app deberia romper/apagarse, NO se usa para errores controlados
		//debemos usar panic solo cuando la app no puede seguir de forma segura o no puede continuar simplemente

		panic(err)
	}

	app.Logger.Println("App is up!")

	//primero registramos todos los handlers, luego escuchamos con ListenAndServe
	//de esta forma bindeamos paths con function handlers en nuestro server
	http.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr:        ":8080",
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatalln("We couldnt spin up our server ")
	}

}

// http handler
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//w: interface usada por HTTP handlers para crear respuestas HTTP
	//	con el contestamos al cliente

	//r: HTTP request recibida por el servidor, lo que nos envia el cliente

	fmt.Fprintln(w, "Server is up and running")
}
