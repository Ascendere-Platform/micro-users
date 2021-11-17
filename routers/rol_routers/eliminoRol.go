package rolrouters

import (
	"net/http"

	rolbd "github.com/ascendere/micro-users/bd/rol_bd"
)

func EliminarRol (w http.ResponseWriter, r *http.Request){

	rolID := r.URL.Query().Get("id")

	if len(rolID)<1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := rolbd.BorroRol(rolID)
	if err != nil {
		http.Error(w, "Ocurrio un error" + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}