package services

import (
	"api/src/commons/banco"
	"api/src/models/envelope"
	"api/src/repository"
)

// BuscarEnvelopes trás os envelopes do banco de dados
func BuscarEnvelopes(userId uint) ([]envelope.Envelope, error) {
	bd, erro := banco.Conectar()

	if erro != nil {
		return nil, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.GetEnvelopes(userId)
}

func BuscaEnvelopePorId(envelopeId, userId uint) (envelope.Envelope, error) {
	bd, erro := banco.Conectar()

	if erro != nil {
		return envelope.Envelope{}, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.GetEnvelopePorId(envelopeId, userId)

}

func BuscarEnvelopePorNome(nome string, userId uint) ([]envelope.Envelope, error) {

	bd, erro := banco.Conectar()

	if erro != nil {
		return nil, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.GetEnvelopePorNome(nome, userId)
}

func InserirNovoEnvelope(envelope envelope.Envelope) (uint, error) {
	bd, erro := banco.Conectar()

	if erro != nil {
		return 0, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.Insert(envelope)
}

func DeletarEnvelopePorID(envelopeId, userId uint) error {
	bd, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.DeleteById(envelopeId, userId)
}

func AtualizarEnvelope(envelope envelope.Envelope, userId uint) error {
	db, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceEnvelope(db)
	return repositorio.AtualizarEnvelope(envelope, userId)
}
