package rotas

import (
	"api/src/controller"
	"net/http"
)

var associacaoDespesaRecorrencia = []Rota{
	{
		URI:          "/associacao-despesa-recorrencia",
		Metodo:       http.MethodPost,
		Func:         controller.SalvaAssociacaoDespesaRecorrencia,
		Autenticacao: true,
	},
	{
		URI:          "/associacao-despesa-recorrencia/desfazer/{recorrenciaId}",
		Metodo:       http.MethodDelete,
		Func:         controller.DesfazerAssociacaoDespesaRecorrencia,
		Autenticacao: true,
	},
}
