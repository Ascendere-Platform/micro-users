package usuariosrouters

import (
	"encoding/json"
	"net/http"

	usuariosbd "github.com/ascendere/micro-users/bd/usuarios_bd"
	"github.com/ascendere/micro-users/models"
	"github.com/ascendere/micro-users/routers"
)

func ModificarPerfil (w http.ResponseWriter, r *http.Request){

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(),400)
		return
	}

	var status bool
	status, err = usuariosbd.ModificoRegistro(t, routers.IDUsuario)

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