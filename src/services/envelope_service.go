package services

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
)

// BuscarEnvelopes trás os envelopes do banco de dados
func BuscarEnvelopes() ([]models.Envelope, error) {
	bd, erro := banco.Conectar()

	if erro != nil {
		return nil, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)

	return repositorio.GetEnvelopes()
}

func BuscarEnvelopePorNome(nome string) (models.Envelope, error) {

	bd, erro := banco.Conectar()

	if erro != nil {
		return models.Envelope{}, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.GetEnvelopePorNome(nome)
}
