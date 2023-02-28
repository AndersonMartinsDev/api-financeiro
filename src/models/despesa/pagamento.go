package despesa

import "api/src/models"

type Pagamento struct {
	Id             uint        `json:"id,omitempty"`
	Valor          float64     `json:"valor"`
	DataPagamento  models.Date `json:"datapagamento,omitempty"`
	DataVencimento models.Date `json:"dataVencimento,omitempty"`
	FormaPagamento string      `json:"formapagamento,omitempty"`
	UsuarioId      uint        `json:"usuarioid,omitempty"`
	DespesaId      uint        `json:"despesaid"`
}
