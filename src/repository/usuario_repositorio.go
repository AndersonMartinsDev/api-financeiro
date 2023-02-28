package repository

import (
	"api/src/models/usuario"
	"database/sql"
)

type UsuarioRepositorio struct {
	sql *sql.DB
}

func NewInstanceUsuario(sql *sql.DB) *UsuarioRepositorio {
	return &UsuarioRepositorio{sql}
}

// NovoUsuario criação de novo usuario
func (repositorio UsuarioRepositorio) NovoUsuario(usuario usuario.Usuario) error {
	insert := `Insert Into usuario(avatar,nome,username,senha,email) values(?,?,?,?,?)`
	statement, erro := repositorio.sql.Prepare(insert)

	if erro != nil {
		return nil
	}

	_, erro = statement.Exec(usuario.Avatar, usuario.Nome, usuario.Username, usuario.Senha, usuario.Email)
	return erro
}

func (repositorio UsuarioRepositorio) UsuarioPorID(usuarioId uint) (usuario.Usuario, error) {
	query := `Select id,avatar,nome,username,email from usuario where id = ?`
	row, erro := repositorio.sql.Query(query, usuarioId)

	if erro != nil {
		return usuario.Usuario{}, erro
	}

	var usuarioE usuario.Usuario
	if row.Next() {
		if erro := row.Scan(
			&usuarioE.ID,
			&usuarioE.Avatar,
			&usuarioE.Nome,
			&usuarioE.Username,
			&usuarioE.Email,
		); erro != nil {
			return usuario.Usuario{}, erro
		}
	}
	return usuarioE, nil
}

func (repositorio UsuarioRepositorio) UsuarioDTOPorID(usuarioId uint) (usuario.UsuarioDTO, error) {
	query := `Select id,avatar,nome,username,email from usuario where id = ?`
	row, erro := repositorio.sql.Query(query, usuarioId)

	if erro != nil {
		return usuario.UsuarioDTO{}, erro
	}

	var entity usuario.UsuarioDTO
	if row.Next() {
		if erro := row.Scan(
			&entity.ID,
			&entity.Avatar,
			&entity.Nome,
		); erro != nil {
			return usuario.UsuarioDTO{}, erro
		}
	}
	return entity, nil
}
