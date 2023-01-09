package controller

import (
	"api/src/models"
	"api/src/respostas"
	"api/src/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SalvaAssociacaoDespesaRecorrencia(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var associacao models.AssociacaoDespesaRecorrencia
	if erro := json.Unmarshal(body, &associacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	erro = services.NovaAssociacao(associacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, "Salvo com Sucesso!")

}

func DesfazerAssociacaoDespesaRecorrencia(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	recorrenciaId, erro := strconv.ParseUint(parametro["recorrenciaId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	erro = services.RemoveAssociao(uint(recorrenciaId))

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, erro)
		return
	}
	_ = services.DeletaRecorrencia(uint(recorrenciaId))
	respostas.JSON(w, http.StatusOK, "Sucesso!")
}
