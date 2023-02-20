package repository

import (
	"api/src/models"
	"database/sql"
)

type PagamentoRepositorio struct {
	sql *sql.DB
}

func NewInstancePagamento(sql *sql.DB) *PagamentoRepositorio {
	return &PagamentoRepositorio{sql}
}

func (repositorio *PagamentoRepositorio) Insert(pagamento models.Pagamento) error {

	insert := `INSERT INTO pagamentos(valor,data_vencimento,despesa_id) values(?,?,?)`

	statement, erro := repositorio.sql.Prepare(insert)

	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(pagamento.Valor, pagamento.DataVencimento.Time, pagamento.DespesaId)

	return erro

}
