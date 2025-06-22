package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/joaquinbian/workout-api-go/internal/app"
	"github.com/joaquinbian/workout-api-go/internal/routes"
)

func main() {

	var port int

	//nos deja pasarle el puerto mediante la flag -port y la guarda en la variable port
	flag.IntVar(&port, "port", 8080, "Server port")
	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {
		//log.Fatal("An error ocurred instanciating the app")
		//la funcion panic va a crashear la app. Lo usamos para situaciones que NO DEBERIAN OCURRIR
		//si algo de esto ocurre, no hay nada que hacer, la app deberia romper/apagarse, NO se usa para errores controlados
		//debemos usar panic solo cuando la app no puede seguir de forma segura o no puede continuar simplemente

		panic(err)
	}

	defer app.DB.Close()

	//primero registramos todos los handlers, luego escuchamos con ListenAndServe
	//de esta forma bindeamos paths con function handlers en nuestro server
	routerHandler := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		Handler:     routerHandler,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
	}

	app.Logger.Printf("App is up and listening on port %d", port)

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatalln("We couldnt spin up our server ")
	}

}
