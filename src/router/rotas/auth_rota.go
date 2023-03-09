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
}
