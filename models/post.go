package models

type Post struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}