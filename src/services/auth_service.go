package services

import (
	"api/src/commons/seguranca"
	"api/src/models/usuario"
)

func Login(usuario *usuario.UsuarioLoginDto) error {
	login, erro := UsuarioLoginDtoPorUsername(usuario.Username)
	if erro != nil {
		return erro
	}
	usuario.ID = login.ID
	return seguranca.VerificarSenha(login.Senha, usuario.Senha)
}
