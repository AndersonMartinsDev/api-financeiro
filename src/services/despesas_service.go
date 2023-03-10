package services

import (
	"api/src/commons/banco"
	"api/src/models/despesa"
	"api/src/repository"
)

// GetDespesas busca todas as despesas do banco
func GetDespesas(carteira string) ([]despesa.VDespesa, error) {
	db, erro := banco.Conectar()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	despesas, erro := repositorio.GetDespesas(carteira)

	if erro != nil {
		return nil, erro
	}
	return despesas, nil
}

// GetDespesaPorId busca despesa por id
func GetDespesasById(despesaId uint, carteira string) (despesa.Despesa, error) {
	db, erro := banco.Conectar()

	if erro != nil {
		return despesa.Despesa{}, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	entity, erro := repositorio.GetDespesasById(despesaId, carteira)
	if erro != nil {
		return despesa.Despesa{}, erro
	}

	return entity, nil
}

// NovaDespesa cria uma nova despesa
func NovaDespesa(entity despesa.DespesaPagamento) (uint, error) {
	db, erro := banco.Conectar()

	if erro != nil {
		return 0, erro
	}
	defer db.Close()

	if erro := entity.Despesa.Check(); erro != nil {
		return 0, erro
	}

	repositorio := repository.NewInstanceDespesa(db)
	id, erro := repositorio.Insert(entity.Despesa)

	go func() {
		if entity.Despesa.Tipo == despesa.PARCELADA {
			for _, v := range entity.Pagamentos {
				v.DespesaId = id
				if erro := v.Check(); erro == nil {
					InserirPagamento(v)
				}
			}

		}
	}()

	if erro != nil {
		return 0, erro
	}
	return id, nil
}

func AtualizaStatusQuitacaoDespesa(despesaId uint, quitada bool) error {
	db, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	return repositorio.UpdateStatusQuitacao(despesaId, quitada)
}

// AtualizaDespesa atualiza os valores de despesa
func AtualizaDespesa(despesa despesa.Despesa) error {
	db, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorioDespesa := repository.NewInstanceDespesa(db)
	erro = repositorioDespesa.Update(despesa)

	if erro != nil {
		return erro
	}
	return nil
}

func AtualizaAssociacaoDespesaEnvelope(despesaId, envelopeId uint) error {
	db, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	return repositorio.AtualizaEnvelopeDespesa(despesaId, envelopeId)
}

func GetTotalDespesaMes(carteira string) (float64, error) {
	db, erro := banco.Conectar()
	if erro != nil {
		return 0, erro
	}
	defer db.Close()
	repositorio := repository.NewInstanceDespesa(db)
	return repositorio.GetTotalDespesaPorMes(carteira)
}

func DeletaDespesa(despesaId uint, carteira string) error {
	db, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	return repositorio.DeletaDespesa(despesaId, carteira)
}
