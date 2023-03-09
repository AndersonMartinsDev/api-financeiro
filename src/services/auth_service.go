package services

import (
	"api/src/commons/seguranca"
	"api/src/models/usuario"
)

func Login(usuario usuario.UsuarioLoginDto) error {
	login, erro := UsuarioLoginPorUsername(usuario.Username)
	if erro != nil {
		return erro
	}
	return seguranca.VerificarSenha(usuario.Senha, login.Senha)
}
