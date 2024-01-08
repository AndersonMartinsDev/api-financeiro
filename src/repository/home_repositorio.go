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

func (repositorio HomeRepositorio) GetTotaisCard(usuarioId uint) (home.TotaisCardInfo, error) {
	query := `with card as (
					select (
							select AVG(md.valor) from v_despesa md  
							where 
								md.usuario = ? and
								md.data_vencimento BETWEEN DATE_SUB(now(),INTERVAL 12 MONTH) 
								and NOW()
							) as media , 
							vd.quitada
							from v_despesa vd  
							where 
							vd.usuario = ? and
							DATE_FORMAT(vd.data_vencimento,'%m/%Y') = DATE_FORMAT(NOW(),'%m/%Y')
				)
				select media,
						CONCAT((select COUNT(quitada)from card where card.quitada=1),'/',COUNT(quitada)) as fracao  
						from card`

	linha, erro := repositorio.sql.Query(query, usuarioId, usuarioId)
	if erro != nil {
		return home.TotaisCardInfo{}, erro
	}
	defer linha.Close()

	var entity home.TotaisCardInfo

	if linha.Next() {
		if erro := linha.Scan(
			&entity.MediaDespesas,
			&entity.FracaoContas,
		); erro != nil {
			return home.TotaisCardInfo{}, erro
		}
	}

	return entity, nil
}
