package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Facultad struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NombreFacultad string `bson:"nombreRol" json:"nombreRol,omitempty"`
} 