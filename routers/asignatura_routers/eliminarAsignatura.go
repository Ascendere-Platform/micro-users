package asignaturarouters

import (
	"net/http"

	asignaturabd "github.com/ascendere/micro-users/bd/asignatura_bd"
)

func EliminarAsignatura(w http.ResponseWriter, r *http.Request) {

	asignaturaID := r.URL.Query().Get("id")

	if len(asignaturaID) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := asignaturabd.BorroAsignatura(asignaturaID)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}