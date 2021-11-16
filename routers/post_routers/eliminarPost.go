package postrouters

import (
	"net/http"

	postbd "github.com/ascendere/micro-users/bd/post_bd"
	"github.com/ascendere/micro-users/routers"
)

func EliminarTweet (w http.ResponseWriter, r *http.Request){

	ID := r.URL.Query().Get("id")

	if len(ID)<1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := postbd.BorroPost(ID, routers.IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error" + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}