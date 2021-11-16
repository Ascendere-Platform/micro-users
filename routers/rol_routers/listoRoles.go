package rolrouters

import (
	"encoding/json"
	"net/http"

	rolbd "github.com/ascendere/micro-users/bd/rol_bd"
)

func ListaRoles(w http.ResponseWriter, r *http.Request) {

	result, status := rolbd.LeoRolesTodos()
	if !status {
		http.Error(w, "Error al leer los roles", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}