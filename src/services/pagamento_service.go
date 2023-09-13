package services

import (
	"api/src/commons/banco"
	"api/src/models/despesa"
	"api/src/repository"
)

func InserirPagamento(pagamentos despesa.Pagamento) error {
	bd, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstancePagamento(bd)
	return repositorio.Insert(pagamentos)
}

func GetPagamentosPorDespesaId(despesaId uint) ([]despesa.PagamentoDto, error) {

	bd, erro := banco.Conectar()
	if erro != nil {
		return nil, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstancePagamento(bd)
	return repositorio.GetPagamentos(despesaId)
}

func IndicarPagamento(pagamento despesa.PagamentoUpdateDto) error {
	bd, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstancePagamento(bd)
	erro = repositorio.IndicarPagamento(pagamento)
	if erro != nil {
		return erro
	}

	lastId, despesaId, erro := repositorio.VerificaUltimoPagamento(pagamento)
	if erro != nil {
		return erro
	}

	if lastId == pagamento.Id {
		AtualizaStatusQuitacaoDespesa(despesaId, true)
	}
	return erro
}

func VerificaUltimoPagamento(pagamento despesa.PagamentoUpdateDto) error {
	bd, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstancePagamento(bd)
	return repositorio.IndicarPagamento(pagamento)
}

func IndicarNovoPagamento(pagamento despesa.Pagamento) error {
	bd, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstancePagamento(bd)
	return repositorio.IndicarNovoPagamento(pagamento)

}

func RemoverPagamento(pagamentoId uint) error {
	bd, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstancePagamento(bd)
	return repositorio.RemoverPagamento(pagamentoId)
}
