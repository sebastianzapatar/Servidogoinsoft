package handlers

import (
	"ingsoft/middlw"
	"ingsoft/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlw.CheckBD(routes.Registro)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3023"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(PORT, handler))
}
