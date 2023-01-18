package services

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
)

// NovaRecorrencia insere um novo dado no banco chamando o repositorio
func NovaRecorrencia(recorrencia models.Recorrencia) (uint, error) {

	db, erro := banco.Conectar()

	if erro != nil {
		return 0, erro
	}
	defer db.Close()

	repositorioRecorrencia := repository.NewInstanceRecorrencia(db)

	repositorioRecorrencia.Insert(recorrencia)

	return uint(recorrencia.Id), nil
}

func DeletaRecorrencia(recorrenciaId uint) error {
	db, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorioRecorrencia := repository.NewInstanceRecorrencia(db)
	return repositorioRecorrencia.DeletaRecorrencia(recorrenciaId)
}
