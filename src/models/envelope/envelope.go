package envelope

type Envelope struct {
	Id         uint    `json:"id,omitempty"`
	Titulo     string  `json:"titulo,omitempty"`
	Valor      float64 `json:"valor,omitempty"`
	Observacao string  `json:"observacao,omitempty"`
	Carteira   string
}
