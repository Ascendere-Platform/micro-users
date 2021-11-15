package models

type UsuariosAsignaturas struct {
	UsuarioID string  `bson:"usuarioid" json:"usuarioId"`
	AsignaturaID string `bson:"asignaturaId" json:"asignaturaId"`
}