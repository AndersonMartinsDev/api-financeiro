package repository

import (
	"api/src/models/despesa"
	"database/sql"
)

var (
	updateQuitacao = `update despesas set quitada=? where id = ?`
)

type DespesaRepositorio struct {
	sql *sql.DB
}

// NewInstanceDespesa cria nova instancia do repositorio
func NewInstanceDespesa(banco *sql.DB) *DespesaRepositorio {
	return &DespesaRepositorio{banco}
}

func (repositorio DespesaRepositorio) GetDespesasById(despesaId uint, userId uint) (despesa.Despesa, error) {
	query := `SELECT 		
				DISTINCT(des.id),
				des.titulo,
				des.valor,
				IF(( Date_format(pgto.data_pagamento, '%m/%y') =
                    Date_format(Now(), '%m/%y')
                       ), true, false) AS quitada,
				des.tipo,
				IF(des.tipo <> 'PARCELADA',des.dia_vencimento,DAY(pgto.data_vencimento)) AS dia_vencimento,
				des.observacao,
				IF(env.id IS NULL,0,env.id),
				IF(env.titulo IS NULL,"",env.titulo),
				IF(env.titulo IS NULL,"",env.observacao)
			from despesas des
			LEFT join envelopes env on env.id = des.envelope_id
			LEFT JOIN pagamentos pgto ON pgto.despesa_id = des.id
			`
	linhas, erro := repositorio.sql.Query(query+" where des.id = ? and des.usuario_id = ? ", despesaId, userId)
	if erro != nil {
		return despesa.Despesa{}, erro
	}
	defer linhas.Close()

	var entity despesa.Despesa

	if linhas.Next() {
		if erro := linhas.Scan(
			&entity.ID,
			&entity.Titulo,
			&entity.Valor,
			&entity.Quitada,
			&entity.Tipo,
			&entity.DiaVencimento,
			&entity.Observacao,
			&entity.Envelope.Id,
			&entity.Envelope.Titulo,
			&entity.Envelope.Observacao,
		); erro != nil {
			return despesa.Despesa{}, erro
		}
	}
	return entity, nil
}

// GetDespesas tras todas as despesas gerais baseada nas despesas cadastradas
func (repositorio DespesaRepositorio) GetDespesas(user_id uint, filter string) ([]despesa.VDespesa, error) {
	//SE ALTERAR ESSA QUERY DEVE ALTERAR A DO BALANÇO TOTAL TAMBÉM
	queryDefault := `AND
					(DATE_FORMAT(d.data_vencimento,'%m/%Y') = DATE_FORMAT(NOW(),'%m/%Y') 
							OR d.tipo = 'FIXA' 
							OR (d.tipo = 'UNICA' 
								AND DATE_FORMAT(d.data_cadastro,'%m/%Y') = DATE_FORMAT(NOW(),'%m/%Y')))`
	queryFilter := ` AND
						((d.tipo = 'FIXA' AND DATE_FORMAT(d.data_cadastro,'%m/%Y') = DATE_FORMAT(?,'%m/%Y'))
						OR (d.tipo = 'UNICA' AND (DATE_FORMAT(?,'%m/%Y') >= DATE_FORMAT(d.data_cadastro,'%m/%Y')) AND data_vencimento IS NOT NULL)
						OR DATE_FORMAT(d.data_vencimento,'%m/%Y') = DATE_FORMAT(?,'%m/%Y'))`

	queryView := `SELECT 
					id,
					titulo,
					valor,
					condicao,
					pagamento,
					quitada 
					FROM v_despesa d 
					WHERE 
					d.usuario = ?
					`
	var linhas *sql.Rows
	var erro error

	if filter != "null" {
		linhas, erro = repositorio.sql.Query(queryView+queryFilter, user_id, filter, filter, filter)
	} else {
		linhas, erro = repositorio.sql.Query(queryView+queryDefault, user_id)
	}

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var despesas []despesa.VDespesa

	for linhas.Next() {
		var despesa despesa.VDespesa
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
func (repositorio DespesaRepositorio) Insert(despesa despesa.Despesa) (uint, error) {
	insert := `Insert into despesas(titulo, valor, quitada, tipo, dia_vencimento, observacao, envelope_id, usuario_id) values(?,?,?,?,?,?,?,?)`
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

	result, erro := statement.Exec(despesa.Titulo, despesa.Valor, despesa.Quitada, despesa.Tipo, despesa.DiaVencimento, despesa.Observacao, envelope, despesa.Usuario)

	if erro != nil {
		return 0, erro
	}

	ID, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint(ID), nil
}
func (repositorio DespesaRepositorio) Update(despesa despesa.Despesa) error {

	update := `update despesas set valor=?, quitada=?, dia_vencimento=?, envelope_id=? where id=?`
	statement, erro := repositorio.sql.Prepare(update)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	var envelope interface{}
	if despesa.Envelope.Id != 0 {
		envelope = despesa.Envelope.Id
	}

	_, erro = statement.Exec(despesa.Valor, despesa.Quitada, despesa.DiaVencimento, envelope, despesa.ID)

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
func (repositorio DespesaRepositorio) GetTotalDespesaPorMes(userId uint) (float64, error) {
	query_total_mes := `SELECT 
							ROUND(SUM(valor) , 2)
							FROM v_despesa d
							WHERE d.usuario = ?
							AND
							(DATE_FORMAT(d.data_vencimento,'%m/%y') = DATE_FORMAT(NOW(),'%m/%y') 
									OR d.tipo = 'FIXA' 
									OR (d.tipo = 'UNICA' 
										AND DATE_FORMAT(d.data_cadastro,'%m/%y') = DATE_FORMAT(NOW(),'%m/%y')))`
	total, erro := repositorio.sql.Query(query_total_mes, userId)
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
func (repositorio DespesaRepositorio) DeletaDespesa(despesaID uint, userId uint) error {

	deletePagamentos := `delete from pagamentos where despesa_id = (select id from despesas where id=? and quitada <> 1);`
	delete := `delete from despesas where id =? and quitada <> 1`

	stmPagemtos, erro := repositorio.sql.Prepare(deletePagamentos)
	if erro != nil {
		return erro
	}
	defer stmPagemtos.Close()

	statement, erro := repositorio.sql.Prepare(delete)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	stmPagemtos.Exec(despesaID)
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
