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
				IF(env.id IS NULL,0,env.id),
				IF(env.titulo IS NULL,"",env.titulo),
				IF(rec.id IS NULL,0,rec.id),
				IF(rec.meses IS NULL,0,rec.meses)
			from despesas des
			left join envelope env on env.id = des.envelope_id
			left join recorrencia rec on rec.id = des.recorrencia_id`
	insert         = `Insert into despesas(titulo, valor, quitada, fixa, dia_vencimento, envelope_id, recorrencia_id) values(?,?,?,?,?,?,?)`
	update         = `update despesas set valor=?, quitada=?, fixa=?, dia_vencimento=?, envelope_id=?,recorrencia_id= ? where id=?`
	delete         = `delete from despesas where id = ?`
	updateQuitacao = `update despesas set quitada=? where id = ?`
)

type DespesaRepositorio struct {
	sql *sql.DB
}

// NewInstanceDespesa cria nova instancia do repositorio
func NewInstanceDespesa(banco *sql.DB) *DespesaRepositorio {
	return &DespesaRepositorio{banco}
}
func (repositorio DespesaRepositorio) GetDespesasByNome(nome string) ([]models.Despesa, error) {
	linhas, erro := repositorio.sql.Query(query+" where des.titulo LIKE ? ", "%"+nome+"%")

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
			&despesa.Recorrencia.Id,
			&despesa.Recorrencia.Meses,
		); erro != nil {
			return nil, erro
		}
		despesas = append(despesas, despesa)
	}
	return despesas, nil

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
			&despesa.Recorrencia.Id,
			&despesa.Recorrencia.Meses,
		); erro != nil {
			return nil, erro
		}
		despesas = append(despesas, despesa)
	}

	return despesas, nil
}

// Insert insere um novo registro de despesa
func (repositorio DespesaRepositorio) Insert(despesa models.Despesa) (uint, error) {

	statement, erro := repositorio.sql.Prepare(insert)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	var recorrencia interface{}
	recorrencia = despesa.Recorrencia.Id
	if despesa.Recorrencia.Id == 0 {
		recorrencia = nil
	}

	var envelope interface{}
	envelope = despesa.Envelope.Id
	if despesa.Envelope.Id == 0 {
		envelope = nil
	}

	result, erro := statement.Exec(despesa.Titulo, despesa.Valor, despesa.Quitada, despesa.Fixa, despesa.DiaVencimento, envelope, recorrencia)

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

	var recorrencia interface{}
	if despesa.Recorrencia.Id != 0 {
		recorrencia = despesa.Recorrencia.Id
	}

	var envelope interface{}
	if despesa.Envelope.Id != 0 {
		envelope = despesa.Envelope.Id
	}

	_, erro = statement.Exec(despesa.Valor, despesa.Quitada, despesa.Fixa, despesa.DiaVencimento, envelope, recorrencia, despesa.ID)

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

func (repositorio DespesaRepositorio) UpdateStatusQuitacao(despesaId uint, quitada bool) error {
	statement, erro := repositorio.sql.Prepare(updateQuitacao)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(quitada, despesaId)
	return erro
}
