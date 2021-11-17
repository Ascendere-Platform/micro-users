package asignaturarouters

import (
	"encoding/json"
	"net/http"

	asignaturabd "github.com/ascendere/micro-users/bd/asignatura_bd"
	"github.com/ascendere/micro-users/models"
)

func IngresarAsignatura(w http.ResponseWriter, r *http.Request) {
	var asig models.Asignatura

	err := json.NewDecoder(r.Body).Decode(&asig)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	_, status, err := asignaturabd.RegistroAsignatura(asig)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar una nueva facultad", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar una nueva facultad", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}