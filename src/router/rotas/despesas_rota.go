package rotas

import (
	"api/src/controller"
	"net/http"
)

var despesasRotas = []Rota{
	{
		URI:          "/despesas",
		Metodo:       http.MethodGet,
		Func:         controller.GetDespesas,
		Autenticacao: true,
	},
	{
		URI:          "/despesas",
		Metodo:       http.MethodPost,
		Func:         controller.NovaDespesa,
		Autenticacao: true,
	},
	{
		URI:          "/despesas/{despesaTitulo}",
		Metodo:       http.MethodGet,
		Func:         controller.GetDespesaPorNome,
		Autenticacao: true,
	},
	{
		URI:          "/despesas",
		Metodo:       http.MethodPut,
		Func:         controller.AtualizaDespesa,
		Autenticacao: true,
	},
	{
		URI:          "/despesas/{despesaId}/{quitada}",
		Metodo:       http.MethodPut,
		Func:         controller.AtualizaQuitacaoDespesa,
		Autenticacao: true,
	},
	{
		URI:          "/despesas/{id}",
		Metodo:       http.MethodDelete,
		Func:         controller.DeletaDespesa,
		Autenticacao: true,
	},
}
