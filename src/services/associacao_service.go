package services

import (
	"api/src/commons/banco"
	"api/src/models/associacao"
	"api/src/repository"
)

func NovaAssociacaoCarteiraUsuario(associacao associacao.AssociacaoCarteiraUsuario) error {
	hashIdCarteira, erro := NovaHashCarteira(associacao.CarteiraId)
	if erro != nil {
		return erro
	}
	associacao.CarteiraId = string(hashIdCarteira)

	bd, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer bd.Close()

	repository := repository.NewInstanceAssociacaoRepositorio(bd)
	return repository.NovaAssociacaoCarteiraUsuario(associacao)
}
