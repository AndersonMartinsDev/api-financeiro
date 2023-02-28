package services

import (
	"api/src/models/usuario"
	"api/src/repository"
	"api/src/tools/banco"
)

// InserirUsuario insere novo usuario
func InserirUsuario(usuario usuario.Usuario) error {
	db, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceUsuario(db)
	return repositorio.NovoUsuario(usuario)
}

// UsuarioPorId busca todos os atributos de usuario exceto a senha
func UsuarioPorId(usuarioId uint) (usuario.Usuario, error) {
	db, erro := banco.Conectar()
	if erro != nil {
		return usuario.Usuario{}, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceUsuario(db)
	return repositorio.UsuarioPorID(usuarioId)
}

// UsuarioDTOid busca uma versão enxuta de usuario
func UsuarioDTOid(usuarioId uint) (usuario.UsuarioDTO, error) {
	db, erro := banco.Conectar()
	if erro != nil {
		return usuario.UsuarioDTO{}, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceUsuario(db)
	return repositorio.UsuarioDTOPorID(usuarioId)
}
