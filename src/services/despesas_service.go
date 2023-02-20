package services

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
)

// GetDespesas busca todas as despesas do banco
func GetDespesas() ([]models.VDespesa, error) {
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
func GetDespesasById(despesaId uint) (models.Despesa, error) {
	db, erro := banco.Conectar()

	if erro != nil {
		return models.Despesa{}, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	despesa, erro := repositorio.GetDespesasById(despesaId)
	if erro != nil {
		return models.Despesa{}, erro
	}

	return despesa, nil
}

// NovaDespesa cria uma nova despesa
func NovaDespesa(despesa models.DespesaPagamento) (uint, error) {
	db, erro := banco.Conectar()

	if erro != nil {
		return 0, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	id, erro := repositorio.Insert(despesa.Despesa)

	go func() {
		if despesa.Despesa.Tipo == models.PARCELADA {
			for _, v := range despesa.Pagamentos {
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
func AtualizaDespesa(despesa models.Despesa) error {
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
