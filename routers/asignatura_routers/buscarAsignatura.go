package asignaturarouters

import (
	"encoding/json"
	"net/http"

	asignaturabd "github.com/ascendere/micro-users/bd/asignatura_bd"
)

//VerPerfil permite extraer los valores del Rol
func BuscarAsignatura (w http.ResponseWriter, r *http.Request) {

	asignatura := r.URL.Query().Get("asignatura")


	informacion, err := asignaturabd.BuscoAsignatura(asignatura)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar una Asignatura ", 400)
		return 
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}