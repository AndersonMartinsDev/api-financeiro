package repository

import (
	"api/src/models/associacao"
	"database/sql"
)

type AssociacaoRepositorio struct {
	*sql.DB
}

func NewInstanceAssociacaoRepositorio(sql *sql.DB) *AssociacaoRepositorio {
	return &AssociacaoRepositorio{sql}
}

func (repositorio AssociacaoRepositorio) NovaAssociacaoCarteiraUsuario(associacao associacao.AssociacaoCarteiraUsuario) error {
	insert := `Insert into associacao_carteira_usuario(usuario_id,carteira_id) values(?,?)`
	statement, erro := repositorio.DB.Prepare(insert)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(associacao.UsuarioId, associacao.CarteiraId)
	return erro
}
