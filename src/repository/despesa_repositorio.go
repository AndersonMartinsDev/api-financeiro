package repository

import (
	"api/src/models"
	"database/sql"
)

var (
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
func (repositorio DespesaRepositorio) GetDespesasById(despesaId uint) (models.Despesa, error) {
	query := `SELECT 		
				DISTINCT(des.id),
				des.titulo,
				des.valor,
				des.quitada,
				des.tipo,
				IF(des.tipo <> 'PARCELADA',des.dia_vencimento,DAY(pgto.data_vencimento)) AS dia_vencimento,
				IF(env.id IS NULL,0,env.id),
				IF(env.titulo IS NULL,"",env.titulo),
				IF(env.titulo IS NULL,"",env.observacao)
			from despesas des
			LEFT join envelopes env on env.id = des.envelope_id
			LEFT JOIN pagamentos pgto ON pgto.despesa_id = des.id
			`
	linhas, erro := repositorio.sql.Query(query+" where des.id = ? ", despesaId)
	if erro != nil {
		return models.Despesa{}, erro
	}
	defer linhas.Close()

	var despesa models.Despesa

	if linhas.Next() {
		if erro := linhas.Scan(
			&despesa.ID,
			&despesa.Titulo,
			&despesa.Valor,
			&despesa.Quitada,
			&despesa.Tipo,
			&despesa.DiaVencimento,
			&despesa.Envelope.Id,
			&despesa.Envelope.Titulo,
			&despesa.Envelope.Observacao,
		); erro != nil {
			return models.Despesa{}, erro
		}
	}
	return despesa, nil

}

// GetDespesas tras todas as despesas gerais baseada nas despesas cadastradas
func (repositorio DespesaRepositorio) GetDespesas() ([]models.VDespesa, error) {
	//SE ALTERAR ESSA QUERY DEVE ALTERAR A DO BALANÇO TOTAL TAMBÉM
	queryView := `SELECT 
					id,
					titulo,
					valor,
					condicao,
					pagamento,
					quitada 
					FROM v_despesa d 
					WHERE DATE_FORMAT(d.data_vencimento,'%m/%y') = DATE_FORMAT(NOW(),'%m/%y') 
							OR d.tipo = 'FIXA' 
							OR (d.tipo = 'UNICA' 
								AND DATE_FORMAT(d.data_cadastro,'%m/%y') = DATE_FORMAT(NOW(),'%m/%y'))`

	linhas, erro := repositorio.sql.Query(queryView)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var despesas []models.VDespesa

	for linhas.Next() {
		var despesa models.VDespesa
		if erro := linhas.Scan(
			&despesa.ID,
			&despesa.Titulo,
			&despesa.Valor,
			&despesa.Condicao,
			&despesa.Pagamento,
			&despesa.Quitada,
		); erro != nil {
			return nil, erro
		}
		despesas = append(despesas, despesa)
	}

	return despesas, nil
}

// Insert insere um novo registro de despesa
func (repositorio DespesaRepositorio) Insert(despesa models.Despesa) (uint, error) {
	insert := `Insert into despesas(titulo, valor, quitada, tipo, dia_vencimento, observacao, envelope_id) values(?,?,?,?,?,?,?)`
	statement, erro := repositorio.sql.Prepare(insert)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	var envelope interface{}
	envelope = despesa.Envelope.Id
	if despesa.Envelope.Id == 0 {
		envelope = nil
	}

	result, erro := statement.Exec(despesa.Titulo, despesa.Valor, despesa.Quitada, despesa.Tipo, despesa.DiaVencimento, despesa.Observacao, envelope)

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

	var envelope interface{}
	if despesa.Envelope.Id != 0 {
		envelope = despesa.Envelope.Id
	}

	_, erro = statement.Exec(despesa.Valor, despesa.Quitada, despesa.Tipo, despesa.DiaVencimento, envelope, despesa.ID)

	if erro != nil {
		return erro
	}
	return nil
}
func (repositorio DespesaRepositorio) AtualizaEnvelopeDespesa(despesaId, envelopeId uint) error {
	updateDespesaEnvelope := `UPDATE despesas des
							 SET envelope_id = ?
							 WHERE des.id = ?`
	statement, erro := repositorio.sql.Prepare(updateDespesaEnvelope)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(envelopeId, despesaId)
	return erro
}

func (repositorio DespesaRepositorio) GetTotalDespesaPorMes() (float64, error) {
	query_total_mes := `SELECT SUM(valor) 
							FROM v_despesa d
							WHERE DATE_FORMAT(d.data_vencimento,'%m/%y') = DATE_FORMAT(NOW(),'%m/%y') 
									OR d.tipo = 'FIXA' 
									OR (d.tipo = 'UNICA' 
										AND DATE_FORMAT(d.data_cadastro,'%m/%y') = DATE_FORMAT(NOW(),'%m/%y'))`
	total, erro := repositorio.sql.Query(query_total_mes)
	if erro != nil {
		return 0, erro
	}

	var totalValor float64
	if total.Next() {
		if erro := total.Scan(
			&totalValor,
		); erro != nil {
			return 0, erro
		}
	}
	return totalValor, erro
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
