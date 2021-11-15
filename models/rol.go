package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rol struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NombreRol string `bson:"nombreRol" json:"nombreRol,omitempty"`
} 