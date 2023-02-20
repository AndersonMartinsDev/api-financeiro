package models

type Pagamento struct {
	Id             uint    `json:"id,omitempty"`
	Valor          float64 `json:"valor"`
	DataPagamento  Date    `json:"datapagamento,omitempty"`
	DataVencimento Date    `json:"dataVencimento,omitempty"`
	FormaPagamento string  `json:"formapagamento,omitempty"`
	UsuarioId      uint    `json:"usuarioid,omitempty"`
	DespesaId      uint    `json:"despesaid"`
}
