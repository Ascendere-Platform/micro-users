package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UsuariosAsignaturas struct {
	UsuarioID    string `bson:"usuarioid" json:"usuarioId"`
	AsignaturaID primitive.ObjectID `bson:"asignaturaId" json:"asignaturaId"`
}