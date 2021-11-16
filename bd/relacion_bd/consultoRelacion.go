package relacionbd

import (
	"context"
	"fmt"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}