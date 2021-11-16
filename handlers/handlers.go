package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ascendere/micro-users/middlew"
	usuariosrouters "github.com/ascendere/micro-users/routers/usuarios_routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(usuariosrouters.Registro)).Methods("POST")


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
