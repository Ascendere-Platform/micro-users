package usuariosAsignaturasrouters

import (
	"net/http"

	usuariosAsignaturasbd "github.com/ascendere/micro-users/bd/usuariosAsignaturas_bd"
	"github.com/ascendere/micro-users/models"
	"github.com/ascendere/micro-users/routers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EliminarAsignaturaUsuario(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")

	if len(ID)<1{
		http.Error(w, "El id es obligatorio", http.StatusBadRequest)
		return
	}

	objID, errConversion := primitive.ObjectIDFromHex(ID)

	if errConversion != nil {
		http.Error(w, "Ocurrio un error al convertir el id", http.StatusBadRequest)
		return
	}

	var t models.UsuariosAsignaturas
	t.UsuarioID = routers.IDUsuario
	t.AsignaturaID = objID

	status, err := usuariosAsignaturasbd.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Error al intentar borrar la relación"+ err.Error(), http.StatusBadRequest)
	}

	if !status {
		http.Error(w, "No se logro borrar la relación" + err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}