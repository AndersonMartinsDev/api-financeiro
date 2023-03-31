package controller

import (
	"api/src/commons/autenticacao"
	"api/src/commons/respostas"
	"api/src/models/usuario"
	"api/src/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario usuario.UsuarioLoginDto
	if erro := json.Unmarshal(body, &usuario); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if erro := services.Login(&usuario); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuario)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, token)
}

func CriarNovaCarteira(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario usuario.UsuarioResumeDTO
	if erro := json.Unmarshal(body, &usuario); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	usuarioBD, _ := services.UsuarioResumeDTOUsername(usuario.Username)
	if erro := services.AssociacaoCarteiraUsuario(usuarioBD.ID); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, true)
}

func ExisteCarteiraVinculada(w http.ResponseWriter, r *http.Request) {
	username, erro := autenticacao.ExtrairUsername(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	temCarteira := services.ExisteCarteiraVinculada(username)
	respostas.JSON(w, http.StatusOK, temCarteira)
}
