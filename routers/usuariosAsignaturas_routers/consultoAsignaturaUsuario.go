package usuariosAsignaturasrouters

import (
	"encoding/json"
	"net/http"

	usuariosAsignaturasbd "github.com/ascendere/micro-users/bd/usuariosAsignaturas_bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConsultarAsignaturaUsuario (w http.ResponseWriter, r *http.Request) {
	usuarioID := r.URL.Query().Get("usuarioId")
	asignaturaID := r.URL.Query().Get("asignaturaId")

	objID, errConversion := primitive.ObjectIDFromHex(asignaturaID)

	if errConversion != nil {
		http.Error(w, "Ocurrio un error al convertir el id", http.StatusBadRequest)
		return
	}

	var t models.UsuariosAsignaturas
	t.UsuarioID = usuarioID
	t.AsignaturaID = objID

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