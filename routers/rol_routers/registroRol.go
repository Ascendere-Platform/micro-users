package rolrouters

import (
	"encoding/json"
	"net/http"

	rolbd "github.com/ascendere/micro-users/bd/rol_bd"
	"github.com/ascendere/micro-users/models"
)

func RegistroRol(w http.ResponseWriter, r *http.Request){
	var rol models.Rol

	err := json.NewDecoder(r.Body).Decode(&rol)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+ err.Error(),400)
		return
	}

	_, econtrado, _ := rolbd.ChequeoYaExisteRol(rol.NombreRol)
	if econtrado {
		http.Error(w, "Ya existe el rol "+rol.NombreRol, 400)
		return
	}

	_, status, err := rolbd.RegistroRol(rol)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar un nuevo rol", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar un nuevo rol", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}