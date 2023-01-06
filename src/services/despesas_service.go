package services

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
)

// GetDespesas busca todas as despesas do banco
func GetDespesas(despesaFiltro models.Despesa) ([]models.Despesa, error) {
	db, erro := banco.Conectar()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	despesas, erro := repositorio.GetDespesas(despesaFiltro)
	print(repositorio)
	return despesas, nil
}

// NovaDespesa cria uma nova despesa
func NovaDespesa(despesa models.Despesa) (uint, error) {
	db, erro := banco.Conectar()

	if erro != nil {
		return 0, erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceDespesa(db)
	id, erro := repositorio.Insert(despesa)

	if erro != nil {
		return 0, erro
	}
	return id, nil
}
