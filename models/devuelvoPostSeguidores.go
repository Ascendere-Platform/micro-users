package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoPostSeguidores struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID string `bson:"usuarioid" json:"userID,omitempty"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"userRelationId,omitempty"`
	Post struct {
		Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID string `bson:"_id" json:"_id,omitempty"`
	}
}