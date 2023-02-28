package despesa

type DespesaPagamento struct {
	Despesa    Despesa     `json:"despesa"`
	Pagamentos []Pagamento `json:"pagamentos"`
}
