package repository

import (
	"api/src/models/home"
	"database/sql"
)

type HomeRepositorio struct {
	sql *sql.DB
}

func NewInstanceHome(banco *sql.DB) *HomeRepositorio {
	return &HomeRepositorio{banco}
}

func (repositorio HomeRepositorio) GetTotais(usuarioId uint) (home.TotaisInfoModel, error) {
	query := `with dash as (SELECT valor,
								quitada 
								FROM v_despesa d 
								WHERE 
								d.usuario = ?
								AND
								(DATE_FORMAT(d.data_vencimento,'%m/%Y') = DATE_FORMAT(NOW(),'%m/%Y') 
										OR d.tipo = 'FIXA' 
										OR (d.tipo = 'UNICA' 
											AND DATE_FORMAT(d.data_cadastro,'%m/%Y') = DATE_FORMAT(NOW(),'%m/%Y'))))
								select sum(valor) as balanco_total,(select sum(valor) from dash where quitada = 1) as total_pago from dash`

	linha, erro := repositorio.sql.Query(query, usuarioId)

	if erro != nil {
		return home.TotaisInfoModel{}, erro
	}

	defer linha.Close()

	var entity home.TotaisInfoModel

	if linha.Next() {
		if erro := linha.Scan(
			&entity.BalancoMes,
			&entity.TotalPago,
		); erro != nil {
			return home.TotaisInfoModel{}, erro
		}
	}

	return entity, nil
}
