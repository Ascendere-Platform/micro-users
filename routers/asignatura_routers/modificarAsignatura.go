package asignaturarouters

import (
	"encoding/json"
	"net/http"

	asignaturabd "github.com/ascendere/micro-users/bd/asignatura_bd"
	"github.com/ascendere/micro-users/models"
)

func ModificarAsignatura (w http.ResponseWriter, r *http.Request){

	var t models.Asignatura

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(),400)
		return
	}

	var status bool
	status, err = asignaturabd.ModificoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro"+err.Error(),400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el registro",400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	
}