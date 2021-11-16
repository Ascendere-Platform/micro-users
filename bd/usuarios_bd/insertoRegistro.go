package usuariosbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoRegistro es la parada final con la BD para insertar los datos del usuario
func InsertoRegistro(u models.Usuario) (string, bool, error){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)
	u.ID = primitive.NewObjectID()
	u.RolId = "6194238cc5b410303ff7b50d"

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}