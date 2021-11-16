package rolbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ChequeoUsuarioAdmin(id string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("usuarios")
	colRol := db.Collection("rol")

	var resultadoUsuario models.Usuario
	var resultadoRol models.Rol

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return resultadoUsuario, false, "No se pudo transformar el id"
	}

	condicion := bson.M{"_id": objID}
	condicionRol := bson.M{"nombreRol": "admin"}

	errUsuario := col.FindOne(ctx, condicion).Decode(&resultadoUsuario)

	errRol := colRol.FindOne(ctx, condicionRol).Decode(&resultadoRol)

	if errUsuario != nil {
		return resultadoUsuario, false, "No se encuentra el usuario"
	}
	if errRol != nil {
		return resultadoUsuario, false, "No existe el rol de ADMINISTRADOR"
	}

	idUsuarioRol, error := primitive.ObjectIDFromHex(resultadoUsuario.RolId)

	if error != nil {
		return resultadoUsuario, false, "Error al encontrar usuario"
	}
	if idUsuarioRol != resultadoRol.ID {
		return resultadoUsuario, false, "No es un ADMINISTRADOR"
	}

	return resultadoUsuario, true, id
}