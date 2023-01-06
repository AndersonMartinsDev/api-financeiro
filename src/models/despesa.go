package models

type Despesa struct {
	ID           uint64   `json:"id,omitempty"`
	Titulo       string   `json:"titulo,omitempty"`
	Valor        float64  `json:"valor,omitempty"`
	Quitada      bool     `json:"quitada,omitempty"`
	DataCadastro string   `json:"datacadastro,omitempty"`
	Envelope     Envelope `json:"envelope,omitempty"`
}
