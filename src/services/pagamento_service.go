package services

import (
	"api/src/models/despesa"
	"api/src/repository"
	"api/src/tools/banco"
)

func InserirPagamento(pagamentos despesa.Pagamento) error {
	bd, erro := banco.Conectar()
	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstancePagamento(bd)
	return repositorio.Insert(pagamentos)
}
