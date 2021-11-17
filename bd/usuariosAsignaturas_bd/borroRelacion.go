package usuariosAsignaturasbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
)

func BorroRelacion (t models.UsuariosAsignaturas) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Usuarios")
	col := db.Collection("usuariosAsignaturas")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}
	
	return true, nil
}