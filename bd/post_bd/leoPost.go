package postbd

import (
	"context"
	"log"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoTweets es la función que me devuelve todos los tweets de un usuario
func LeoTweets (ID string, pagina int64) ([]*models.DevuelvoPost, bool){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("post")

	var resultados [] * models.DevuelvoPost

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key:"fecha", Value: -1}})
	opciones.SetSkip((pagina -1)*20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()){
		var registro models.DevuelvoPost
		err := cursor.Decode(&registro)

		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}

	return resultados, true

}