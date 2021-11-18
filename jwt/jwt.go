package jwt

import (
	"time"

	"github.com/ascendere/micro-users/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GeneroJWT(t models.Usuario) (string, error){

	miClave := []byte("MastersDelUniverso")

	payload := jwt.MapClaims{
		"email": t.Email,
		"nombre": t.Nombres,
		"apellidos": t.Apellidos,
		"rolId": t.RolId,
		"fecha_Nacimiento": t.FechaNacimiento,
		"biografia": t.Biografia,
		"ubicacion": t.Ubicacion,
		"sitioweb": t.SitioWeb,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 168).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}