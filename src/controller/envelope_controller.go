package controller

import (
	"api/src/respostas"
	"api/src/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// InsereNovoEnvelope busca o servico para inserir o dado no banco
func InsereNovoEnvelope(w http.ResponseWriter, r *http.Request) {

}

// BuscarEnvelopes busca o servico para inserir o dado no banco
func BuscarEnvelopes(w http.ResponseWriter, r *http.Request) {
	response, erro := services.BuscarEnvelopes()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, response)
}

// BuscarEnvelopePorNome busca o servico para inserir o dado no banco
func BuscarEnvelopePorNome(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	nome := parametro["envelopeId"]
	envelope, erro := services.BuscarEnvelopePorNome(nome)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, envelope)
}

// BuscaEnvelopePorId busca o servico para inserir o dado no banco
func BuscaEnvelopePorId(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	envelopeId, erro := strconv.ParseUint(parametro["envelopeId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	fmt.Print(envelopeId)

}

// DeletaEnvelopePorId busca o servico para inserir o dado no banco
func DeletaEnvelopePorId(w http.ResponseWriter, r *http.Request) {

}

// AtualizaEnvelope busca o servico para inserir o dado no banco
func AtualizaEnvelope(w http.ResponseWriter, r *http.Request) {

}
