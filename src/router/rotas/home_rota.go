package rotas

import (
	"api/src/controller"
	"net/http"
)

var homeRotas = []Rota{
	{
		URI:          "/home/totais",
		Metodo:       http.MethodGet,
		Func:         controller.GetTotaisChart,
		Autenticacao: true,
	},
	{
		URI:          "/home/cards",
		Metodo:       http.MethodGet,
		Func:         controller.GetTotaisCard,
		Autenticacao: true,
	},
}
