package rotas

import (
	"api/src/controller"
	"net/http"
)

var usuarioRouter = []Rota{
	{
		URI:          "/usuario",
		Metodo:       http.MethodPost,
		Func:         controller.NovoUsuario,
		Autenticacao: true,
	},
	{
		URI:          "/usuario/{usuarioId}",
		Metodo:       http.MethodGet,
		Func:         controller.UsuarioPorId,
		Autenticacao: true,
	},
	{
		URI:          "/usuario/{usuarioId}/dto",
		Metodo:       http.MethodGet,
		Func:         controller.UsuarioDTOPorId,
		Autenticacao: true,
	},
}
