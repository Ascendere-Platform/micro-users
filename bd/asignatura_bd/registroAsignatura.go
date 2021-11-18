package asignaturabd

import (
	"context"
	"strings"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistroAsignatura (a models.Asignatura) (string, bool, error){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("asignatura")

	modalidad := strings.ToUpper(a.Modalidad)
	periodo := strings.ToUpper(a.Periodo)

	registro := models.Asignatura{
		ID:        primitive.NewObjectID(),
		NombreAsignatura: a.NombreAsignatura,
		Modalidad: modalidad,
		Periodo: periodo,
		FacultadID: a.FacultadID,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}