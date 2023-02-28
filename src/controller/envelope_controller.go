package controller

import (
	"api/src/models/envelope"
	"api/src/services"
	"api/src/tools/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// InsereNovoEnvelope busca o servico para inserir o dado no banco
func InsereNovoEnvelope(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var entity envelope.Envelope
	if erro := json.Unmarshal(body, &entity); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response, erro := services.InserirNovoEnvelope(entity)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, response)
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
	nome := parametro["nome"]
	envelopes, erro := services.BuscarEnvelopePorNome(nome)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, envelopes)
}

// BuscaEnvelopePorId busca o servico para inserir o dado no banco
func BuscaEnvelopePorId(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	envelopeId, erro := strconv.ParseUint(parametro["envelopeId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	envelope, erro := services.BuscaEnvelopePorId(uint(envelopeId))
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, envelope)
}

// DeletaEnvelopePorId busca o servico para inserir o dado no banco
func DeletaEnvelopePorId(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	envelopeId, erro := strconv.ParseUint(parametro["envelopeId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if erro := services.DeletarEnvelopePorID(uint(envelopeId)); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, "Deletado com sucesso")
}

// AtualizaEnvelope busca o servico para inserir o dado no banco
func AtualizaEnvelope(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var envelope envelope.Envelope
	if erro := json.Unmarshal(body, &envelope); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if erro := services.AtualizarEnvelope(envelope); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, "Atualizado com sucesso")
}
