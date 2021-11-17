package routers

import (
	"encoding/json"
	"net/http"
)

func Home (w http.ResponseWriter, r *http.Request) {


	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Acceso Denegado")

}