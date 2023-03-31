package rotas

import (
	"api/src/controller"
	"net/http"
)

var authRoutes = []Rota{
	{
		URI:          "/login",
		Metodo:       http.MethodPost,
		Func:         controller.Login,
		Autenticacao: false,
	},
	{
		URI:          "/carteira/nova",
		Metodo:       http.MethodPost,
		Func:         controller.CriarNovaCarteira,
		Autenticacao: true,
	},
	{
		URI:          "/carteira/vinculo",
		Metodo:       http.MethodPost,
		Func:         controller.CriarNovaCarteira,
		Autenticacao: true,
	},
	{
		URI:          "/carteira/existe-vinculo",
		Metodo:       http.MethodGet,
		Func:         controller.ExisteCarteiraVinculada,
		Autenticacao: true,
	},
}
