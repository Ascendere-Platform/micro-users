package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
//Usuario es el modelo de usuarios de la BD
type DevuelvoUsuario struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Rol Rol `bson:"rol" json:"rol,omitempty"`
	Nombres string `bson:"nombre" json:"nombre,omitempty"`
	FechaNacimiento time.Time `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email string `bson:"email" json:"email"`
	Avatar string `bson:"avatar" json:"avatar,omitempty"`
	Banner string `bson:"banner" json:"banner,omitempty"`
	Biografia string `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion string `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb string `bson:"sitioWeb" json:"sitioWeb,omitempty"`
}