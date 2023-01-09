package repository

import (
	"api/src/models"
	"database/sql"
)

type AssociacaoDespesaRecorrenciaRepository struct {
	sql *sql.DB
}

func NewInstanceAssociacao(db *sql.DB) *AssociacaoDespesaRecorrenciaRepository {
	return &AssociacaoDespesaRecorrenciaRepository{db}
}

// Insert insere uma nova associação entre as despesas e recorrencias
func (repositorio AssociacaoDespesaRecorrenciaRepository) Insert(associacao models.AssociacaoDespesaRecorrencia) {
	statement, erro := repositorio.sql.Prepare("Insert into associacao_despesa_recorrencia(despesa_id,recorrencia_id) values (?,?)")

	if erro != nil {
		return
	}
	defer statement.Close()

	resultado, _ := statement.Exec(associacao.Despesa.ID, associacao.Recorrencia.Id)

	print(resultado.LastInsertId())
}

func (repositorio AssociacaoDespesaRecorrenciaRepository) RemoveAssociao(recorrenciaId uint) error {
	statement, erro := repositorio.sql.Prepare("Delete from associacao_despesa_recorrencia where recorrencia_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(recorrenciaId)
	return erro
}
