package relacionrouters

import (
	"encoding/json"
	"net/http"

	relacionbd "github.com/ascendere/micro-users/bd/relacion_bd"
	"github.com/ascendere/micro-users/models"
	"github.com/ascendere/micro-users/routers"
)

func ConsultoRelacion (w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = routers.IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := relacionbd.ConsultoRelacion(t)

	if err != nil || !status{
		resp.Status=false
	} else {
		resp.Status=true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)
}