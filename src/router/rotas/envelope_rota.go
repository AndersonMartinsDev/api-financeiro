package rotas

import (
	"api/src/controller"
	"net/http"
)

var envelopeRotas = []Rota{
	{
		URI:          "/envelopes",
		Metodo:       http.MethodGet,
		Func:         controller.BuscarEnvelopes,
		Autenticacao: true,
	},
	{
		URI:          "/envelopes/{nome}/nomes",
		Metodo:       http.MethodGet,
		Func:         controller.BuscarEnvelopePorNome,
		Autenticacao: true,
	},
	{
		URI:          "/envelopes/{envelopeId}",
		Metodo:       http.MethodGet,
		Func:         controller.BuscaEnvelopePorId,
		Autenticacao: true,
	},
	{
		URI:          "/envelopes",
		Metodo:       http.MethodPost,
		Func:         controller.InsereNovoEnvelope,
		Autenticacao: true,
	},
	{
		URI:          "/envelopes",
		Metodo:       http.MethodPut,
		Func:         controller.AtualizaEnvelope,
		Autenticacao: true,
	},
	{
		URI:          "/envelopes/{envelopeId}",
		Metodo:       http.MethodDelete,
		Func:         controller.DeletaEnvelopePorId,
		Autenticacao: true,
	},
}
