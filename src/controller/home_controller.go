package controller

import (
	"api/src/commons/autenticacao"
	"api/src/commons/respostas"
	"api/src/services"
	"net/http"
)

func GetTotaisChart(w http.ResponseWriter, r *http.Request) {
	user, erro := autenticacao.ExtrairUsername(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	totais, erro := services.GetTotaisChart(user)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, totais)
}
