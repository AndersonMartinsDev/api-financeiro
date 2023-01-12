package repository

import (
	"api/src/models"
	"database/sql"
)

type EnvelopeRepositorio struct {
	sql *sql.DB
}

func NewInstanceEnvelope(sql *sql.DB) *EnvelopeRepositorio {
	return &EnvelopeRepositorio{sql}
}

func (repository EnvelopeRepositorio) GetEnvelopes() ([]models.Envelope, error) {

	linhas, erro := repository.sql.Query("Select id, titulo, valor from envelope")

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var envelopes []models.Envelope
	for linhas.Next() {
		var envelope models.Envelope
		erro = _inserirValorAModel(linhas, &envelope)
		if erro == nil {
			envelopes = append(envelopes, envelope)
		}
	}
	return envelopes, erro
}

func (repository EnvelopeRepositorio) GetEnvelopePorNome(nome string) (models.Envelope, error) {
	linha, erro := repository.sql.Query("Select id,titulo, valor from envelope where nome=?", nome)
	if erro != nil {
		return models.Envelope{}, erro
	}
	defer linha.Close()

	var envelope models.Envelope
	if linha.Next() {
		erro = _inserirValorAModel(linha, &envelope)
	}

	return envelope, erro
}

func _inserirValorAModel(linha *sql.Rows, envelope *models.Envelope) error {
	if erro := linha.Scan(
		&envelope.Id,
		&envelope.Titulo,
		&envelope.Valor,
	); erro != nil {
		return erro
	}
	return nil
}
