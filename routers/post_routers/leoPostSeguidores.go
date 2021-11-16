package postrouters

import (
	"encoding/json"
	"net/http"
	"strconv"

	postbd "github.com/ascendere/micro-users/bd/post_bd"
	"github.com/ascendere/micro-users/routers"
)

/*LeoPostSeguidores lee los tweets de todos nuestros seguidores */
func LeoPostSeguidores (w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := postbd.LeoPostSeguidores(routers.IDUsuario, pagina)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}