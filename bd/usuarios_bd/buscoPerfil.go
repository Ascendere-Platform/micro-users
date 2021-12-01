package usuariosbd

import (
	"context"
	"fmt"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(ID string) (models.DevuelvoUsuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	var usuarioEncontrado models.DevuelvoUsuario
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)

	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return usuarioEncontrado, err
	}

	colRol := db.Collection("rol")

	var rol models.Rol
	objRol,_ := primitive.ObjectIDFromHex(perfil.RolId)

	errRol := colRol.FindOne(ctx, bson.M{"_id": objRol}).Decode(&rol)
	
	if errRol != nil {
		fmt.Println("Rol no encontrado" + errRol.Error())
		return usuarioEncontrado, errRol
	}

	usuarioEncontrado.ID = perfil.ID
	usuarioEncontrado.Rol = rol
	usuarioEncontrado.Nombres = perfil.Nombres + " " + perfil.Apellidos
	usuarioEncontrado.FechaNacimiento = perfil.FechaNacimiento
	usuarioEncontrado.Email = perfil.Email
	usuarioEncontrado.Avatar = perfil.Avatar
	usuarioEncontrado.Banner = perfil.Banner
	usuarioEncontrado.Biografia = perfil.Biografia
	usuarioEncontrado.Ubicacion = perfil.Ubicacion
	usuarioEncontrado.SitioWeb = perfil.SitioWeb

	return usuarioEncontrado, err
}