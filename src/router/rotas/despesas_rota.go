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
		URI:          "/despesas/total",
		Metodo:       http.MethodGet,
		Func:         controller.GetTotalDespesas,
		Autenticacao: true,
	},
	{
		URI:          "/despesas",
		Metodo:       http.MethodPost,
		Func:         controller.NovaDespesa,
		Autenticacao: true,
	},
	{
		URI:          "/despesas/{id}",
		Metodo:       http.MethodGet,
		Func:         controller.GetDespesasById,
		Autenticacao: true,
	},
	{
		URI:          "/despesas",
		Metodo:       http.MethodPut,
		Func:         controller.AtualizaDespesa,
		Autenticacao: true,
	},
	{
		URI:          "/despesas/{id}",
		Metodo:       http.MethodDelete,
		Func:         controller.DeletaDespesa,
		Autenticacao: true,
	},
	{
		URI:          "/despesas/{id}/{envelopeId}",
		Metodo:       http.MethodPut,
		Func:         controller.AtualizaEnvelopeDespesa,
		Autenticacao: true,
	},
}
