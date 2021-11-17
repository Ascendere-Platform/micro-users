package asignaturabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteFacultad(facultad string) (models.Facultad, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("facultad")

	condicion := bson.M{"nombreFacultad":facultad}

	var resultado models.Facultad

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}