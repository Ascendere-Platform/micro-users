package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoAsignaturas struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID string `bson:"usuarioid" json:"userID,omitempty"`
	UsuariosAsignaturasID string `bson:"usuarioasignaturaid" json:"usuarioasignaturaid,omitempty"`
	Asignatura struct {
		ID string `bson:"_id,omitempty" json:"id"`
		NombreAsignatura string `bson:"nombreAsignatura" json:"nombreAsignatura,omitempty"`
		Modalidad string `bson:"modalidad" json:"modalidad,omitempty"`
		Periodo string `bson:"periodo" json:"periodo,omitempty"`
	}
}