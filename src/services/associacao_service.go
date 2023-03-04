package services

import (
	"api/src/repository"
	"api/src/tools/banco"
)

func NovaAssociacaoCarteiraUsuario(usuarioId uint) error {
	hashIdCarteira, erro := NovaHashCarteira(usuarioId)
	if erro != nil {
		return erro
	}

	bd, erro := banco.Conectar()

	if erro != nil {
		return erro
	}

	defer bd.Close()

	repository := repository.NewInstanceAssociacaoRepositorio(bd)
	return repository.NovaAssociacaoCarteiraUsuario(usuarioId, hashIdCarteira)
}

func NovaAssociacaoCarteiraDespesa(despesaId uint, carteiraiD []byte) error {
	bd, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceAssociacaoRepositorio(bd)
	return repositorio.NovaAssociacaoCarteiraDespesa(despesaId, carteiraiD)

}
