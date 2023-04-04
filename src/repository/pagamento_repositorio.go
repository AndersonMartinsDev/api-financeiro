package repository

import (
	"api/src/models/despesa"
	"database/sql"
)

type PagamentoRepositorio struct {
	sql *sql.DB
}

func NewInstancePagamento(sql *sql.DB) *PagamentoRepositorio {
	return &PagamentoRepositorio{sql}
}

func (repositorio *PagamentoRepositorio) Insert(pagamento despesa.Pagamento) error {

	insert := `INSERT INTO pagamentos(valor,data_vencimento,despesa_id) values(?,?,?)`

	statement, erro := repositorio.sql.Prepare(insert)

	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(pagamento.Valor, pagamento.DataVencimento.Time, pagamento.DespesaId)

	return erro

}

func (repositorio *PagamentoRepositorio) GetPagamentos(despesaId uint) ([]despesa.PagamentoDto, error) {
	query := `Select 
				id, 
				ROUND(valor,2), 
				data_pagamento, 
				data_vencimento, 
				forma_pagamento 
			from pagamentos 
			where despesa_id = ?`

	rows, erro := repositorio.sql.Query(query, despesaId)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var pagamentos []despesa.PagamentoDto
	for rows.Next() {
		var pagamento despesa.PagamentoDto
		if erro := rows.Scan(
			&pagamento.Id,
			&pagamento.Valor,
			&pagamento.DataPagamento,
			&pagamento.DataVencimento.Time,
			&pagamento.FormaPagamento,
		); erro != nil {
			return nil, erro
		}
		pagamentos = append(pagamentos, pagamento)
	}
	return pagamentos, nil
}

func (repositorio *PagamentoRepositorio) IndicarPagamento(pagamento despesa.PagamentoUpdateDto) error {
	atualizar := `Update pagamentos set data_pagamento=NOW(), forma_pagamento=? where id=?`

	update, erro := repositorio.sql.Prepare(atualizar)
	if erro != nil {
		return erro
	}
	defer update.Close()

	_, erro = update.Exec(pagamento.FormaPagamento, pagamento.Id)
	return erro
}

func (repositorio *PagamentoRepositorio) IndicarNovoPagamento(pagamento despesa.Pagamento) error {
	atualizar := `INSERT INTO pagamentos(valor,data_vencimento,data_pagamento,forma_pagamento,despesa_id) values(ROUND(?,2),?,NOW(),?,?)`

	update, erro := repositorio.sql.Prepare(atualizar)
	if erro != nil {
		return erro
	}
	defer update.Close()

	_, erro = update.Exec(pagamento.Valor, pagamento.DataVencimento.Time, pagamento.FormaPagamento, pagamento.DespesaId)
	return erro
}

func (repositorio *PagamentoRepositorio) VerificaUltimoPagamento(pagamento despesa.PagamentoUpdateDto) (uint, uint, error) {
	query := `Select MAX(ptos.id),despesa_id from pagamentos ptos where despesa_id = (Select despesa_id from pagamentos where id = ?)`
	row, erro := repositorio.sql.Query(query, pagamento.Id)
	if erro != nil {
		return 0, 0, erro
	}
	defer row.Close()

	var maxParcelaId uint
	var despesaId uint
	if row.Next() {
		if erro := row.Scan(
			&maxParcelaId,
			&despesaId,
		); erro != nil {
			return 0, 0, nil
		}
	}

	return maxParcelaId, despesaId, nil
}
