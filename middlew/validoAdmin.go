package middlew

import (
	"net/http"

	rolbd "github.com/ascendere/micro-users/bd/rol_bd"
	"github.com/ascendere/micro-users/routers"
)

func ValidoAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, econtrado, mensaje := rolbd.ChequeoUsuarioAdmin(routers.IDUsuario)
		if !econtrado {
			http.Error(w, mensaje, 400)
			return
		}
		next.ServeHTTP(w, r)
	}
}
