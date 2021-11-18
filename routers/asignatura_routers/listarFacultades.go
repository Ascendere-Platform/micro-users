package asignaturarouters

import (
	"encoding/json"
	"net/http"

	asignaturabd "github.com/ascendere/micro-users/bd/asignatura_bd"
)

func ListarFacutlades(w http.ResponseWriter, r *http.Request) {

	result, status := asignaturabd.ListoFacultades()
	if !status {
		http.Error(w, "Error al leer las facultades", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}