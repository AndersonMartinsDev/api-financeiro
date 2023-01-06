package models

type Envelope struct {
	Id     uint    `json:"id,omitempty"`
	Titulo string  `json:"titulo,omitempty"`
	Valor  float64 `json:"valor,omitempty"`
}
