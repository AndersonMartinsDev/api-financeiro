package services

import (
	"api/src/commons/banco"
	"api/src/models/envelope"
	"api/src/repository"
)

// BuscarEnvelopes trás os envelopes do banco de dados
func BuscarEnvelopes(carteira string) ([]envelope.Envelope, error) {
	bd, erro := banco.Conectar()

	if erro != nil {
		return nil, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.GetEnvelopes()
}

func BuscaEnvelopePorId(envelopeId uint, carteira string) (envelope.Envelope, error) {
	bd, erro := banco.Conectar()

	if erro != nil {
		return envelope.Envelope{}, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.GetEnvelopePorId(envelopeId)

}

func BuscarEnvelopePorNome(nome, carteira string) ([]envelope.Envelope, error) {

	bd, erro := banco.Conectar()

	if erro != nil {
		return nil, erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.GetEnvelopePorNome(nome)
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

func DeletarEnvelopePorID(envelopeId uint, carteira string) error {
	bd, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer bd.Close()

	repositorio := repository.NewInstanceEnvelope(bd)
	return repositorio.DeleteById(envelopeId)
}

func AtualizarEnvelope(envelope envelope.Envelope, carteira string) error {
	db, erro := banco.Conectar()

	if erro != nil {
		return erro
	}
	defer db.Close()

	repositorio := repository.NewInstanceEnvelope(db)
	return repositorio.AtualizarEnvelope(envelope)
}
