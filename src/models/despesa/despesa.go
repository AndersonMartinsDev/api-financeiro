package despesa

import (
	"api/src/models/envelope"
	"time"
)

type Despesa struct {
	ID            uint64            `json:"id,omitempty"`
	Titulo        string            `json:"titulo,omitempty"`
	Valor         float64           `json:"valor,omitempty"`
	Quitada       bool              `json:"quitada,omitempty"`
	Tipo          string            `json:"tipo,omitempty"`
	DataCadastro  time.Time         `json:"datacadastro,omitempty"`
	DiaVencimento uint              `json:"diavencimento,omitempty"`
	Observacao    string            `json:"observacao,omitempty"`
	Envelope      envelope.Envelope `json:"envelope,omitempty"`
	Carteira      string
}
