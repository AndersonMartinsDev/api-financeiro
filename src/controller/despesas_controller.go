package controller

import (
	"api/src/models"
	"api/src/respostas"
	"api/src/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetDespesasGerais busca todas as despesas gerais contendo o mes e o ano
func GetDespesas(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var despesaFiltro models.Despesa
	if erro := json.Unmarshal(body, &despesaFiltro); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	despesas, erro := services.GetDespesas(despesaFiltro)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, despesas)
}

// NovaDespesa endpoint responsável por receber uma nova despesa e cadastrar
func NovaDespesa(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var despesa models.Despesa
	if erro := json.Unmarshal(body, &despesa); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	id, erro := services.NovaDespesa(despesa)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, id)
}
