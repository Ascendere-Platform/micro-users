package usuariosAsignaturasrouters

import (
	"encoding/json"
	"net/http"

	usuariosAsignaturasbd "github.com/ascendere/micro-users/bd/usuariosAsignaturas_bd"
	"github.com/ascendere/micro-users/models"
)

func ConsultarAsignaturaUsuario (w http.ResponseWriter, r *http.Request) {
	usuarioID := r.URL.Query().Get("usuarioId")
	asignaturaID := r.URL.Query().Get("asignaturaId")

	var t models.UsuariosAsignaturas
	t.UsuarioID = usuarioID
	t.AsignaturaID = asignaturaID

	var resp models.RespuestaConsultaAsignatura

	status, err := usuariosAsignaturasbd.ConsultoRelacion(t)

	if err != nil || !status{
		resp.Status=false
	} else {
		resp.Status=true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)
}