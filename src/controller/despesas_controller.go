package controller

import (
	"api/src/commons/autenticacao"
	"api/src/commons/respostas"
	"api/src/models/despesa"
	"api/src/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AtualizaEnvelopeDespesa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	envelopeId, _ := strconv.ParseUint(parametros["envelopeId"], 10, 64)
	despesaId, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	erro = services.AtualizaEnvelope(uint(despesaId), uint(envelopeId))
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, "Despesa recebeu novo envelope!")
}

// GetDespesasGerais busca todas as despesas gerais contendo o mes e o ano
func GetDespesas(w http.ResponseWriter, r *http.Request) {
	carteira, erro := autenticacao.ExtrairCarteiraId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	lista, erro := services.GetDespesas(carteira)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, lista)
}

// GetTotalDespesas é soma de todas as despesas do mês
func GetTotalDespesas(w http.ResponseWriter, r *http.Request) {
	carteira, erro := autenticacao.ExtrairCarteiraId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	row, erro := services.GetTotalDespesaMes(carteira)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, row)
}

// NovaDespesa endpoint responsável por receber uma nova despesa e cadastrar
func NovaDespesa(w http.ResponseWriter, r *http.Request) {
	carteira, erro := autenticacao.ExtrairCarteiraId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var despesaPagamento despesa.DespesaPagamento
	if erro := json.Unmarshal(body, &despesaPagamento); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	despesaPagamento.Despesa.Carteira = carteira
	id, erro := services.NovaDespesa(despesaPagamento)
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

	var despesa despesa.Despesa
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

// GetDespesasById devolve uma despesa
func GetDespesasById(w http.ResponseWriter, r *http.Request) {
	carteira, erro := autenticacao.ExtrairCarteiraId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	parametro := mux.Vars(r)

	despesaId, erro := strconv.ParseUint(parametro["id"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	despesas, erro := services.GetDespesasById(uint(despesaId), carteira)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, despesas)
}

// DeletaDespesa remove o registro da despesa da base
func DeletaDespesa(w http.ResponseWriter, r *http.Request) {
	carteira, erro := autenticacao.ExtrairCarteiraId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	parametro := mux.Vars(r)

	despesaId, erro := strconv.ParseUint(parametro["id"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	erro = services.DeletaDespesa(uint(despesaId), carteira)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, "Deletado com sucesso!")
}
