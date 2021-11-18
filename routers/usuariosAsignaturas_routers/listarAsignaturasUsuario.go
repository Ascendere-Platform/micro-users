package usuariosAsignaturasrouters

import (
	"encoding/json"
	"net/http"
	"strconv"

	usuariosAsignaturasbd "github.com/ascendere/micro-users/bd/usuariosAsignaturas_bd"
	"github.com/ascendere/micro-users/routers"
)

func ListarAsignaturasUsuario(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := usuariosAsignaturasbd.LeoAsignaturasUsuario(routers.IDUsuario, pagina)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}