package postrouters

import (
	"encoding/json"
	"net/http"
	"time"

	postbd "github.com/ascendere/micro-users/bd/post_bd"
	"github.com/ascendere/micro-users/models"
	"github.com/ascendere/micro-users/routers"
)

func GraboPost (w http.ResponseWriter, r *http.Request){
	var mensaje models.Post

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboPost{
		UserId: routers.IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha: time.Now(),
	}

	_, status, err := postbd.InsertoPost(registro)

	if  err != nil {
		http.Error(w, "Ocurrio un error al insertar el post"+err.Error(), 400)
		return
	}

	if !status{
		http.Error(w, "No se ha logrado insertar el post", 400)
	}

	w.WriteHeader(http.StatusCreated)
}