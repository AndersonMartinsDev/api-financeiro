package services

import (
	"api/src/repository"
	"api/src/tools/banco"

	"golang.org/x/crypto/bcrypt"
)

func NovaHashCarteira(usuarioID uint) ([]byte, error) {

	bd, erro := banco.Conectar()
	if erro != nil {
		return nil, erro
	}

	defer bd.Close()

	usuarioRepositorio := repository.NewInstanceUsuario(bd)
	usuario, erro := usuarioRepositorio.UsuarioPorID(usuarioID)

	if erro != nil {
		return nil, erro
	}

	hash := usuario.Username + usuario.Nome
	return bcrypt.GenerateFromPassword([]byte(hash), bcrypt.DefaultCost)

}
