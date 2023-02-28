package services

import (
	"api/src/models/despesa"
	"api/src/repository"
	"api/src/tools/banco"
)

// GetDespesas busca todas as despesas do banco
func GetDespesas() ([]despesa.VDespesa, error) {
	db, erro := banco.Conectar()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	despesas, erro := repositorio.GetDespesas()

	if erro != nil {
		return nil, erro
	}
	return despesas, nil
}

// GetDespesaPorId busca despesa por id
func GetDespesasById(despesaId uint) (despesa.Despesa, error) {
	db, erro := banco.Conectar()

	if erro != nil {
		return despesa.Despesa{}, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	entity, erro := repositorio.GetDespesasById(despesaId)
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

	repositorio := repository.NewInstanceDespesa(db)
	id, erro := repositorio.Insert(entity.Despesa)

	go func() {
		if entity.Despesa.Tipo == despesa.PARCELADA {
			for _, v := range entity.Pagamentos {
				v.DespesaId = id
				InserirPagamento(v)
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

func AtualizaEnvelope(despesaId, envelopeId uint) error {
	db, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	return repositorio.AtualizaEnvelopeDespesa(despesaId, envelopeId)
}

func GetTotalDespesaMes() (float64, error) {
	db, erro := banco.Conectar()
	if erro != nil {
		return 0, erro
	}
	defer db.Close()
	repositorio := repository.NewInstanceDespesa(db)
	return repositorio.GetTotalDespesaPorMes()
}

func DeletaDespesa(despesaId uint) error {
	db, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	return repositorio.DeletaDespesa(despesaId)
}
