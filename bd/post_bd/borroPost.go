package postbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroPost(ID string, UserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("post")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":objID,
		"userid":UserId,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}