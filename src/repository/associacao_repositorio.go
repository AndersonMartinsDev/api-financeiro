package repository

import (
	"database/sql"
)

type AssociacaoRepositorio struct {
	*sql.DB
}

func NewInstanceAssociacaoRepositorio(sql *sql.DB) *AssociacaoRepositorio {
	return &AssociacaoRepositorio{sql}
}

func (repositorio AssociacaoRepositorio) NovaAssociacaoCarteiraUsuario(usuarioID uint, carteiraID []byte) error {
	insert := `Insert into associacao_carteira_usuario(usuario_id,carteira_id) values(?,?)`
	statement, erro := repositorio.DB.Prepare(insert)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(usuarioID, carteiraID)
	return erro
}

func (repositorio AssociacaoRepositorio) NovaAssociacaoCarteiraDespesa(despesaId uint, carteiraId []byte) error {
	insert := `Insert into associacao_carteira_despesa(despesa_id, carteira_id) values(?,?)`

	statement, erro := repositorio.DB.Prepare(insert)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(despesaId, carteiraId)
	return erro
}
