package repository

import (
	"api/src/models"
	"database/sql"
)

type DespesaRepositorio struct {
	sql *sql.DB
}

// NewInstanceDespesa cria nova instancia do repositorio
func NewInstanceDespesa(banco *sql.DB) *DespesaRepositorio {
	return &DespesaRepositorio{banco}
}

// GetDespesas tras todas as despesas gerais baseada nas despesas cadastradas
func (repositorio DespesaRepositorio) GetDespesas(despesaFiltro models.Despesa) ([]models.Despesa, error) {
	linhas, erro := repositorio.sql.Query("Select id,titulo,valor,quitada,data_cadastro,envelope_id from despesas")

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var despesas []models.Despesa
	for linhas.Next() {
		var despesa models.Despesa
		if erro := linhas.Scan(
			&despesa.ID,
			&despesa.Titulo,
			&despesa.Valor,
			&despesa.Quitada,
			&despesa.DataCadastro,
			&despesa.Envelope.Id,
		); erro != nil {
			return nil, erro
		}
		despesas = append(despesas, despesa)
	}

	return despesas, nil
}

// Insert insere um novo registro de despesa
func (repositorio DespesaRepositorio) Insert(despesa models.Despesa) (uint, error) {
	statement, erro := repositorio.sql.Prepare("Insert into despesa values(?,?,?,?,?,?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(despesa.ID, despesa.Titulo, despesa.Valor, despesa.Quitada, despesa.DataCadastro, despesa.Envelope.Id)

	if erro != nil {
		return 0, erro
	}

	ID, erro := result.LastInsertId()

	return uint(ID), nil
}
