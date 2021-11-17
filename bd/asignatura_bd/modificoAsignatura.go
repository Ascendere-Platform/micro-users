package asignaturabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ModificoRegistro(u models.Asignatura) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("asignatura")

	registro := make(map[string]interface{})

	if len(u.NombreAsignatura) > 0{
		registro["nombreAsignatura"] = u.NombreAsignatura
	}
	if len(u.Modalidad) > 0 {
		registro["modalidad"] = u.Modalidad
	}
	if len(u.FacultadID) > 0 {
		registro["apellidos"] = u.FacultadID
	}
	if len(u.Periodo) > 0 {
		registro["avatar"] = u.Periodo
	}

	updtString := bson.M{
		"$set": registro,
	}

	filtro := bson.M{"_id": bson.M{"$eq": u.ID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)

	if err != nil {
		return false, err
	}

	return true, nil

}