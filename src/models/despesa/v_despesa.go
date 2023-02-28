package despesa

type VDespesa struct {
	ID        uint    `json:"id,omitempty"`
	Titulo    string  `json:"titulo,omitempty"`
	Tipo      string  `json:"tipo,omitempty"`
	Valor     float64 `json:"valor,omitempty"`
	Condicao  string  `json:"condicao,omitempty"`
	Pagamento string  `json:"pagamento,omitempty"`
	Quitada   bool    `json:"quitada"`
}
