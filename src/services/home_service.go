package services

import (
	"api/src/commons/banco"
	"api/src/models/home"
	"api/src/repository"
)

func GetTotaisChart(usuarioId uint) (home.TotaisInfoModel, error) {
	db, erro := banco.Conectar()
	if erro != nil {
		return home.TotaisInfoModel{}, erro
	}
	defer db.Close()
	repositorio := repository.NewInstanceHome(db)
	totais, erro := repositorio.GetTotais(usuarioId)

	porcentagem := totais.TotalPago * 100 / totais.BalancoMes
	totais.PorcentagemPie = uint(porcentagem)
	return totais, erro
}
