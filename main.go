package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"chattigo.com/exercises/service"

)

func main() {
 
	r := routes()
  
	log.Fatal(http.ListenAndServe(":8080", r))  // Inicio servidor en puerto 8080 y paso router r
}

func routes() http.Handler {
	r := chi.NewRouter()  
	r.Use(middleware.Logger) // Log solicitudes http
	r.Get("/palyndrome/{number}", service.Palyndrome) //Ruta para ejercicio Palindromos
	r.Get("/digits/{number}", service.Digits)  // Ruta par ejercicio 3 digitos
		
	return r
}




