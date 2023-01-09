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
	statement, erro := repositorio.sql.Prepare("Insert into despesas(titulo,valor,quitada,fixa,envelope_id) values(?,?,?,?,?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(despesa.Titulo, despesa.Valor, despesa.Quitada, despesa.Fixa, despesa.Envelope.Id)

	if erro != nil {
		return 0, erro
	}

	ID, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint(ID), nil
}

func (repositorio DespesaRepositorio) Update(despesa models.Despesa) error {
	statement, erro := repositorio.sql.Prepare("update despesas set valor=?, quitada=?, fixa=?, envelope_id=? where id=?")

	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(despesa.Valor, despesa.Quitada, despesa.Fixa, despesa.Envelope.Id, despesa.ID)

	if erro != nil {
		return erro
	}
	return nil
}

func (repositorio DespesaRepositorio) GetById(despesaId uint) (models.Despesa, error) {
	linha, erro := repositorio.sql.Query(`Select 
												d.id,
												d.titulo,
												d.valor,
												d.quitada,
												d.fixa,
												d.data_cadastro, 
												en.id
												from despesas d
												inner join envelope en on en.id = d.envelope_id
												where d.id = ?
												`, despesaId)
	if erro != nil {
		return models.Despesa{}, erro
	}
	defer linha.Close()

	var despesa models.Despesa
	if linha.Next() {
		if erro := linha.Scan(
			&despesa.ID,
			&despesa.Titulo,
			&despesa.Valor,
			&despesa.Quitada,
			&despesa.Fixa,
			&despesa.DataCadastro,
			&despesa.Envelope.Id,
		); erro != nil {
			return models.Despesa{}, erro
		}
	}
	return despesa, erro
}

func (repositorio DespesaRepositorio) DeletaDespesa(despesaID uint) error {
	statement, erro := repositorio.sql.Prepare("delete from despesas where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(despesaID)
	return erro
}
