package usuariosrouters

import (
	"net/http"

	rolbd "github.com/ascendere/micro-users/bd/rol_bd"
	"github.com/ascendere/micro-users/models"
)

func ModificarRolUsuario (w http.ResponseWriter, r *http.Request){

	idUser := r.URL.Query().Get("iduser")
	idRol := r.URL.Query().Get("idrol")

	var t models.Usuario

	t.RolId = idRol


	var status bool
	status, err := rolbd.ModificoRolUsuario(t, idUser)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el rol del usuario "+err.Error(),400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el rol del usuario ",400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	
}