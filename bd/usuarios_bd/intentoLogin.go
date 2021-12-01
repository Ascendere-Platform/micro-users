package usuariosbd

import (
	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.DevuelvoUsuario, bool) {
	var usuarioEncontrado models.DevuelvoUsuario
	usu, encontrado, _ := bd.ChequeoYaExisteUsuario(email)

	if !encontrado {
		return usuarioEncontrado, false
	}

	usuarioEncontrado, _ = BuscoPerfil(usu.ID.Hex())

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usuarioEncontrado, false
	}

	return usuarioEncontrado, true

}
