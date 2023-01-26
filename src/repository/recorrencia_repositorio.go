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

func (repositorio RecorrenciaRepositorio) Update(recorrencia models.Recorrencia) error {
	statement, erro := repositorio.sql.Prepare("update from recorrencia set meses=? where id=?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(recorrencia.Meses, recorrencia.Id)

	if erro != nil {
		return erro
	}

	return nil

}

func (repositorio RecorrenciaRepositorio) Insert(recorrencia models.Recorrencia) (uint, error) {
	statement, erro := repositorio.sql.Prepare("Insert into recorrencia(meses) values(?)")

	if erro != nil {
		return 0, nil
	}
	defer statement.Close()

	result, erro := statement.Exec(recorrencia.Meses)

	if erro != nil {
		return 0, erro
	}

	recorrenciaId, _ := result.LastInsertId()

	return uint(recorrenciaId), nil
}

func (repositorio RecorrenciaRepositorio) GetByDespesaId(despesaId uint) (models.Recorrencia, error) {
	linha, erro := repositorio.sql.Query("select re.id,re.meses from despesas des inner join recorrencia re on re.id = des.recorrencia_id where des.id =?", despesaId)
	if erro != nil {
		return models.Recorrencia{}, erro
	}
	defer linha.Close()

	var recorrencia models.Recorrencia
	if linha.Next() {
		if erro := linha.Scan(
			&recorrencia.Id,
			&recorrencia.Meses,
		); erro != nil {
			return models.Recorrencia{}, erro
		}
	}
	return recorrencia, nil
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
