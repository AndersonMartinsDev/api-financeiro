package services

import (
	"api/src/commons/banco"
	"api/src/models/associacao"
	"api/src/models/usuario"
	"api/src/repository"
)

// InserirUsuario insere novo usuario
func InserirUsuario(usuario usuario.Usuario) error {
	db, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer db.Close()

	if erro := usuario.Check(); erro != nil {
		return erro
	}

	repositorio := repository.NewInstanceUsuario(db)
	_, erro = repositorio.NovoUsuario(usuario)

	if erro != nil {
		return erro
	}

	return nil
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

// UsuarioLoginPorUsername busca todos os atributos de usuario exceto a senha
func UsuarioLoginPorUsername(username string) (usuario.UsuarioLoginDto, error) {
	db, erro := banco.Conectar()
	if erro != nil {
		return usuario.UsuarioLoginDto{}, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceUsuario(db)
	return repositorio.UsuarioLoginPorUsername(username)
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

// AssociacaoCarteiraUsuario criar uma conexão entre o usuário e a carteira correspondente
func AssociacaoCarteiraUsuario(usuarioId uint) error {

	db, erro := banco.Conectar()
	if erro != nil {
		return nil
	}
	defer db.Close()

	repositorio := repository.NewInstanceUsuario(db)
	usuario, erro := repositorio.UsuarioPorID(usuarioId)

	if erro != nil {
		return nil
	}

	if erro := NovaAssociacaoCarteiraUsuario(associacao.AssociacaoCarteiraUsuario{
		UsuarioId:  usuario.ID,
		CarteiraId: usuario.Username + usuario.Email,
	}); erro != nil {
		return erro
	}

	return nil
}
