package relacionrouters

import (
	"net/http"

	relacionbd "github.com/ascendere/micro-users/bd/relacion_bd"
	"github.com/ascendere/micro-users/models"
	"github.com/ascendere/micro-users/routers"
)

func SeguirUsuario (w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")

	if len(ID)<1{
		http.Error(w, "El id es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = routers.IDUsuario
	t.UsuarioRelacionID = ID

	status, err := relacionbd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error"+ err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se logro insertar" + err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}