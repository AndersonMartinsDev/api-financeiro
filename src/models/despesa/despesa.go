package despesa

import (
	"api/src/models/envelope"
	"fmt"
	"strings"
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
	Usuario       uint
}

func (d *Despesa) Check() error {
	if erro := d.validatedDespesa(); erro != nil {
		return erro
	}
	return nil
}

func (d *Despesa) validatedDespesa() error {

	if erro := validatedText(d.Titulo, "titulo"); erro != nil {
		return erro
	}

	if erro := validatedValue(d.Valor, "titulo"); erro != nil {
		return erro
	}

	if d.Tipo == "" {
		return fmt.Errorf("o campo %s não pode ser vazio", "tipo")
	}

	d.Titulo = strings.TrimSpace(d.Titulo)
	return nil
}

func validatedText(valor, campo string) error {
	if valor == "" {
		return fmt.Errorf("o campo %s não pode ser vazio", campo)
	}
	return nil
}

func validatedValue(valor float64, campo string) error {
	if valor == 0 {
		return fmt.Errorf("o campo %s não pode ser vazio", campo)
	}
	return nil
}
