package models

import "time"

type Despesa struct {
	ID            uint64      `json:"id,omitempty"`
	Titulo        string      `json:"titulo,omitempty"`
	Valor         float64     `json:"valor,omitempty"`
	Quitada       bool        `json:"quitada,omitempty"`
	Fixa          bool        `json:"fixa,omitempty"`
	DataCadastro  time.Time   `json:"datacadastro,omitempty"`
	DiaVencimento uint        `json:"diaVencimento,omitempty"`
	Envelope      Envelope    `json:"envelope,omitempty"`
	Recorrencia   Recorrencia `json:"recorrencia,omitempty"`
}
