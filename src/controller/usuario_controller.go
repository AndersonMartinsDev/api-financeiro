package controller

import (
	"api/src/commons/respostas"
	"api/src/models/usuario"
	"api/src/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// NovoUsuario rota para inserção de novo usuario
func NovoUsuario(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var entity usuario.Usuario
	if erro := json.Unmarshal(body, &entity); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	erro = services.InserirUsuario(entity)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, "Salvo com Sucesso!")
}

// UsuarioPorId busca usuario
func UsuarioPorId(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametro["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	entity, erro := services.UsuarioPorId(uint(usuarioId))

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, entity)
}

// UsuarioDTOPorId busca usuario DTO
func UsuarioDTOPorId(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametro["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	entity, erro := services.UsuarioDTOid(uint(usuarioId))
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, entity)
}
