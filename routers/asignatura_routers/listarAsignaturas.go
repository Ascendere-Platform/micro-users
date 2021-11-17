package asignaturarouters

import (
	"encoding/json"
	"net/http"

	asignaturabd "github.com/ascendere/micro-users/bd/asignatura_bd"
)

func ListarAsignaturas (w http.ResponseWriter, r *http.Request) {

	result, status := asignaturabd.ListoAsignaturas()
	if !status {
		http.Error(w, "Error al leer las asignaturas", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}