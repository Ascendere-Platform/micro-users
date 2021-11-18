package asignaturabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
)


func BuscoFacultad(nombre string) (models.Facultad, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("facultad")

	condicion := bson.M{"nombreFacultad":nombre}

	var resultado models.Facultad

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	
	if err != nil {
		return resultado, err
	}
	return resultado, err
}