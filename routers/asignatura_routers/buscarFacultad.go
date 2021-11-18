package asignaturarouters

import (
	"encoding/json"
	"net/http"

	asignaturabd "github.com/ascendere/micro-users/bd/asignatura_bd"
)

//VerPerfil permite extraer los valores del Rol
func BuscarFacultad (w http.ResponseWriter, r *http.Request) {

	facultad := r.URL.Query().Get("facultad")


	informacion, err := asignaturabd.BuscoFacultad(facultad)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar una Facultad ", 400)
		return 
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}