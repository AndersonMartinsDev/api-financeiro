package despesa

import (
	"api/src/models"
	"database/sql"
)

type PagamentoDto struct {
	Id             uint           `json:"id,omitempty"`
	Valor          float64        `json:"valor"`
	DataPagamento  sql.NullTime   `json:"datapagamento,omitempty"`
	DataVencimento models.Date    `db:"data_vencimento" json:"dataVencimento,omitempty"`
	FormaPagamento sql.NullString `json:"formapagamento,omitempty"`
	DespesaId      uint           `json:"despesaid,omitempty"`
}
