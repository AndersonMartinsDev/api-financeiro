package repository

import (
	"api/src/models"
	"database/sql"
)

type RecorrenciaRepositorio struct {
	sql *sql.DB
}

func NewInstanceRecorrencia(bd *sql.DB) *RecorrenciaRepositorio {
	return &RecorrenciaRepositorio{bd}
}

func (repositorio RecorrenciaRepositorio) Insert(recorrencia models.Recorrencia) (uint, error) {
	statement, erro := repositorio.sql.Prepare("Insert into recorrencia(meses,dia_vencimento) values(?,?)")

	if erro != nil {
		return 0, nil
	}
	defer statement.Close()

	result, erro := statement.Exec(recorrencia.Meses, recorrencia.DiaVencimento)

	if erro != nil {
		return 0, erro
	}

	recorrenciaId, _ := result.LastInsertId()

	return uint(recorrenciaId), nil
}

func (repositorio RecorrenciaRepositorio) DeletaRecorrencia(recorrenciaId uint) error {
	statement, erro := repositorio.sql.Prepare("delete from recorrencia where id = ?")

	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(recorrenciaId)

	return erro
}
