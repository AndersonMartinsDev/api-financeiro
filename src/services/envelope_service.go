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

func BuscaEnvelopePorId(envelopeId uint) (models.Envelope, error) {
	bd, erro := banco.Conectar()

	if erro != nil {
		return models.Envelope{}, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.GetEnvelopePorId(envelopeId)

}

func BuscarEnvelopePorNome(nome string) ([]models.Envelope, error) {

	bd, erro := banco.Conectar()

	if erro != nil {
		return nil, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.GetEnvelopePorNome(nome)
}

func InserirNovoEnvelope(envelope models.Envelope) (uint, error) {
	bd, erro := banco.Conectar()

	if erro != nil {
		return 0, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.Insert(envelope)
}

func DeletarEnvelopePorID(envelopeId uint) error {
	bd, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.DeleteById(envelopeId)
}

func AtualizarEnvelope(envelope models.Envelope) error {
	db, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceEnvelope(db)
	return repositorio.AtualizarEnvelope(envelope)
}
