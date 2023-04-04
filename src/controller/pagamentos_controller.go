package controller

import (
	"api/src/commons/respostas"
	"api/src/models/despesa"
	"api/src/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPagamentos(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	despesaId, erro := strconv.ParseUint(parametro["despesaId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	pagamentos, erro := services.GetPagamentosPorDespesaId(uint(despesaId))
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, pagamentos)
}

func IndicarPagamentos(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var pagamento despesa.PagamentoUpdateDto
	if erro := json.Unmarshal(body, &pagamento); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	erro = services.IndicarPagamento(pagamento)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, "Conta indicada como paga")
}

func IndicarNovoPagamentos(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var pagamento despesa.Pagamento
	if erro := json.Unmarshal(body, &pagamento); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	erro = services.IndicarNovoPagamento(pagamento)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	services.AtualizaStatusQuitacaoDespesa(pagamento.DespesaId, true)

	respostas.JSON(w, http.StatusOK, "Conta nova indicada como paga")
}
