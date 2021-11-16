package rolbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoRolesTodos () ([]*models.Rol, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("rol")

	var results []*models.Rol

	query := bson.M{}

	cur, err := col.Find(ctx, query)
	if err != nil {
		return results, false
	}

	for cur.Next(ctx) {
		var s models.Rol
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}
			results = append(results, &s)

	}

	err = cur.Err()
	if err != nil {
		return results, false
	}
	cur.Close(ctx)
	return results, true
}