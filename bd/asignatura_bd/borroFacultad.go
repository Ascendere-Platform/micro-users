package asignaturabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroFacultad (facultadId string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("facultad")

	objID, _ := primitive.ObjectIDFromHex(facultadId)

	condicion := bson.M{
		"_id":objID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}