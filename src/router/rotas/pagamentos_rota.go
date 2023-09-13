package rotas

import (
	"api/src/controller"
	"net/http"
)

var pagamentosRotas = []Rota{
	{
		URI:          "/pagamentos/{despesaId}",
		Metodo:       http.MethodGet,
		Func:         controller.GetPagamentos,
		Autenticacao: true,
	},
	{
		URI:          "/pagamentos/indicar-pagamento",
		Metodo:       http.MethodPut,
		Func:         controller.IndicarPagamentos,
		Autenticacao: true,
	},
	{
		URI:          "/pagamentos/indicar-novo-pagamento",
		Metodo:       http.MethodPost,
		Func:         controller.IndicarNovoPagamentos,
		Autenticacao: true,
	},
	{
		URI:          "/pagamentos/{pagamentoId}/remover-pagamento",
		Metodo:       http.MethodPut,
		Func:         controller.RemovePagamento,
		Autenticacao: true,
	},
}
