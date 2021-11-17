package asignaturarouters

import (
	"encoding/json"
	"net/http"

	asignaturabd "github.com/ascendere/micro-users/bd/asignatura_bd"
	"github.com/ascendere/micro-users/models"
)

func IngresarFacultad(w http.ResponseWriter, r *http.Request) {
	var fac models.Facultad

	err := json.NewDecoder(r.Body).Decode(&fac)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	_, econtrado, _ := asignaturabd.ChequeoYaExisteFacultad(fac.NombreFacultad)
	if econtrado {
		http.Error(w, "Ya existe la Facultad "+ fac.NombreFacultad, 400)
		return
	}

	_, status, err := asignaturabd.RegistroFacultad(fac)

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