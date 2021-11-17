package usuariosAsignaturasbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoAsignaturasUsuario(idUsuario string, pagina int) ([]models.DevuelvoAsignaturas,bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("usuariosAsignaturas")

	skip := (pagina - 1) *20

	condiciones := make([]bson.M, 0)

	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid":idUsuario}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from": "asignatura",
			"localField": "asignaturaid",
			"foreignField": "_id",
			"as": "asignatura",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$asignatura"})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit":20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoAsignaturas

	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true

}