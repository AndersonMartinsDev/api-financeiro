package services

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
)

func NovaAssociacao(associacao models.AssociacaoDespesaRecorrencia) error {
	bd, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorioDespesa := repository.NewInstanceDespesa(bd)
	despesaId, erro := repositorioDespesa.Insert(associacao.Despesa)

	if erro != nil {
		return erro
	}

	repositorioRecorrencia := repository.NewInstanceRecorrencia(bd)
	recorrenciaId, erro := repositorioRecorrencia.Insert(associacao.Recorrencia)

	if erro != nil {
		return erro
	}
	associacao.Despesa.ID = uint64(despesaId)
	associacao.Recorrencia.Id = uint(recorrenciaId)
	repositorio := repository.NewInstanceAssociacao(bd)
	repositorio.Insert(associacao)

	return nil
}

func RemoveAssociao(recorrenciaId uint) error {
	bd, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceAssociacao(bd)
	return repositorio.RemoveAssociao(recorrenciaId)
}
