package rolbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
)


func BuscoRol(nombre string) (models.Rol, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("rol")

	condicion := bson.M{"nombreRol":nombre}

	var resultado models.Rol

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	
	if err != nil {
		return resultado, err
	}
	return resultado, err
}