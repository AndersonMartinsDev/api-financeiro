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

// AtualizaDespesa atualiza o registro da despesa
func AtualizaDespesa(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var despesa models.Despesa
	if erro := json.Unmarshal(body, &despesa); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	erro = services.AtualizaDespesa(despesa)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, "Atualizado com sucesso!")
}

// GetDespesaPorId devolve uma despesa
func GetDespesaPorId(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	despesaId, erro := strconv.ParseUint(parametro["id"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	despesa, erro := services.GetDespesaPorId(uint(despesaId))
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, despesa)
}

// DeletaDespesa remove o registro da despesa da base
func DeletaDespesa(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	despesaId, erro := strconv.ParseUint(parametro["id"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	erro = services.DeletaDespesa(uint(despesaId))
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, "Deletado com sucesso!")
}
