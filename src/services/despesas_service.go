package services

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
)

// GetDespesas busca todas as despesas do banco
func GetDespesas() ([]models.Despesa, error) {
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
func GetDespesaPorNome(nome string) ([]models.Despesa, error) {
	db, erro := banco.Conectar()

	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	despesas, erro := repositorio.GetDespesasByNome(nome)
	if erro != nil {
		return nil, erro
	}

	return despesas, nil
}

// NovaDespesa cria uma nova despesa
func NovaDespesa(despesa models.Despesa) (uint, error) {
	db, erro := banco.Conectar()

	if erro != nil {
		return 0, erro
	}
	defer db.Close()

	var recorrencia models.Recorrencia

	if despesa.Recorrencia != recorrencia {
		recorrenciaRepositorio := repository.NewInstanceRecorrencia(db)
		recorrenciaId, erro := recorrenciaRepositorio.Insert(despesa.Recorrencia)

		if erro != nil {
			return 0, erro
		}
		despesa.Recorrencia.Id = int64(recorrenciaId)
	}

	repositorio := repository.NewInstanceDespesa(db)
	id, erro := repositorio.Insert(despesa)

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

	repositorio := repository.NewInstanceRecorrencia(db)

	if despesa.Recorrencia.Id != 0 {
		repositorio.Update(despesa.Recorrencia)
	} else {
		if despesa.Recorrencia.Meses > 0 {
			recorrenciaId, _ := repositorio.Insert(despesa.Recorrencia)
			despesa.Recorrencia.Id = int64(recorrenciaId)
		}
	}

	repositorioDespesa := repository.NewInstanceDespesa(db)
	erro = repositorioDespesa.Update(despesa)

	if erro != nil {
		return erro
	}
	return nil
}

func DeletaDespesa(despesaId uint) error {
	db, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)

	erro = repositorio.DeletaDespesa(despesaId)
	if erro != nil {
		return erro
	}

	return nil
}
