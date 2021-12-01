package usuariosrouters

import (
	"encoding/json"
	"net/http"

	usuariosbd "github.com/ascendere/micro-users/bd/usuarios_bd"
)

//VerPerfil permite extraer los valores del Perfil
func VerPerfil (w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := usuariosbd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar un registro "+ err.Error(), 400)
		return 
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}