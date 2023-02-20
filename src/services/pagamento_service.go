package services

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
)

func InserirPagamento(pagamentos models.Pagamento) error {
	bd, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstancePagamento(bd)
	return repositorio.Insert(pagamentos)
}
