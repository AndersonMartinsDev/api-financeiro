package repository

import (
	"api/src/models"
	"database/sql"
)

var (
	query = `Select 
				des.id,
				des.titulo,
				des.valor,
				des.quitada,
				des.fixa,
				des.dia_vencimento,
				des.data_cadastro,
				env.id,
				env.titulo
			from despesas des
			inner join envelope env on env.id = des.envelope_id`
	insert                   = `Insert into despesas(titulo,valor,quitada,fixa,envelope_id,recorrencia_id) values(?,?,?,?,?,?)`
	insertWithoutRecorrencia = `Insert into despesas(titulo,valor,quitada,fixa,envelope_id) values(?,?,?,?,?)`
	update                   = `update despesas set valor=?, quitada=?, fixa=?, envelope_id=? where id=?`
	delete                   = `delete from despesas where id = ?`
)

type DespesaRepositorio struct {
	sql *sql.DB
}

// NewInstanceDespesa cria nova instancia do repositorio
func NewInstanceDespesa(banco *sql.DB) *DespesaRepositorio {
	return &DespesaRepositorio{banco}
}

// GetDespesas tras todas as despesas gerais baseada nas despesas cadastradas
func (repositorio DespesaRepositorio) GetDespesas() ([]models.Despesa, error) {
	linhas, erro := repositorio.sql.Query(query)

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
			&despesa.Fixa,
			&despesa.DiaVencimento,
			&despesa.DataCadastro,
			&despesa.Envelope.Id,
			&despesa.Envelope.Titulo,
		); erro != nil {
			return nil, erro
		}
		despesas = append(despesas, despesa)
	}

	return despesas, nil
}

// Insert insere um novo registro de despesa
func (repositorio DespesaRepositorio) Insert(despesa models.Despesa) (uint, error) {

	insertQuery := insert
	if uint(despesa.Recorrencia.Id) == 0 {
		insertQuery = insertWithoutRecorrencia
	}

	statement, erro := repositorio.sql.Prepare(insertQuery)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	var result sql.Result
	if uint(despesa.Recorrencia.Id) == 0 {
		result, erro = statement.Exec(despesa.Titulo, despesa.Valor, despesa.Quitada, despesa.Fixa, despesa.Envelope.Id)
	} else {
		result, erro = statement.Exec(despesa.Titulo, despesa.Valor, despesa.Quitada, despesa.Fixa, despesa.Envelope.Id, despesa.Recorrencia.Id)
	}

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
	statement, erro := repositorio.sql.Prepare(update)

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
	linha, erro := repositorio.sql.Query(query+" where des.id = ?", despesaId)

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
			&despesa.DiaVencimento,
			&despesa.DataCadastro,
			&despesa.Envelope.Id,
			&despesa.Envelope.Titulo,
		); erro != nil {
			return models.Despesa{}, erro
		}
	}
	return despesa, erro
}

func (repositorio DespesaRepositorio) DeletaDespesa(despesaID uint) error {
	statement, erro := repositorio.sql.Prepare(delete)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(despesaID)
	return erro
}
