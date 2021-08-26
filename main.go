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
  
	log.Fatal(http.ListenAndServe(":8080", r))
}

func routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/palyndrome/{number}", service.Palyndrome)
	r.Get("/digits/{number}", service.Digits)
		
	return r
}




