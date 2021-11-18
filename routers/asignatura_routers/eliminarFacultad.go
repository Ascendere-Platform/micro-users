package asignaturarouters

import (
	"net/http"

	asignaturabd "github.com/ascendere/micro-users/bd/asignatura_bd"
)

func EliminarFacultad(w http.ResponseWriter, r *http.Request) {

	facultadID := r.URL.Query().Get("id")

	if len(facultadID) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := asignaturabd.BorroFacultad(facultadID)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}