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
		URI:          "/despesas/unidade/",
		Metodo:       http.MethodPut,
		Func:         controller.NovaDespesa,
		Autenticacao: true,
	},
}
