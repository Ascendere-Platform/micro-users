package usuariosbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	relacionbd "github.com/ascendere/micro-users/bd/relacion_bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.DevuelvoUsuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("usuarios")

	var results []*models.Usuario
	var resultadoCompleto []*models.DevuelvoUsuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return resultadoCompleto, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			return resultadoCompleto, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false

		encontrado, err = relacionbd.ConsultoRelacion(r)
		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}

		if len(tipo) == 0 {
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""

			results = append(results, &s)
		}
	}

	for _, usuario := range results {
		usuarioEncontrado, _ := BuscoPerfil(usuario.ID.Hex())

		usuarioEncontrado.Biografia = ""
		usuarioEncontrado.SitioWeb = ""
		usuarioEncontrado.Ubicacion = ""
		usuarioEncontrado.Banner = ""

		resultadoCompleto = append(resultadoCompleto, &usuarioEncontrado)
		
	}

	err = cur.Err()
	if err != nil {
		return resultadoCompleto, false
	}
	cur.Close(ctx)
	return resultadoCompleto, true
}