package despesa

import (
	"api/src/models"
	"fmt"
)

type Pagamento struct {
	Id             uint        `json:"id,omitempty"`
	Valor          float64     `json:"valor"`
	DataPagamento  models.Date `json:"datapagamento,omitempty"`
	DataVencimento models.Date `json:"dataVencimento,omitempty"`
	FormaPagamento string      `json:"formapagamento,omitempty"`
	UsuarioId      uint        `json:"usuarioid,omitempty"`
	DespesaId      uint        `json:"despesaid"`
}

func (p *Pagamento) Check() error {
	if erro := p.validateCheck(); erro != nil {
		return erro
	}
	return nil
}

func (p *Pagamento) validateCheck() error {
	if p.DespesaId == 0 {
		return fmt.Errorf("o campo %s não pode ser vazio", "FK despesa")
	}

	date := models.Date{}
	if p.DataVencimento == date {
		return fmt.Errorf("o campo %s não pode ser vazio", "sem data de pagamento")
	}

	return nil
}
