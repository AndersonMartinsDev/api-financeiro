package despesa

type PagamentoUpdateDto struct {
	Id             uint   `json:"id,omitempty"`
	FormaPagamento string `json:"formapagamento,omitempty"`
}
