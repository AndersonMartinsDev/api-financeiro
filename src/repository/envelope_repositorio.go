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

	linhas, erro := repository.sql.Query("Select id, titulo, valor,observacao from envelopes")

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

func (repository EnvelopeRepositorio) GetEnvelopePorNome(nome string) ([]models.Envelope, error) {
	linhas, erro := repository.sql.Query("Select id, titulo, valor, observacao from envelopes where titulo LIKE ?", "%"+nome+"%")
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var envelopes []models.Envelope
	for linhas.Next() {
		var envelope models.Envelope
		if erro := linhas.Scan(
			&envelope.Id,
			&envelope.Titulo,
			&envelope.Valor,
			&envelope.Observacao,
		); erro != nil {
			return nil, erro
		}
		envelopes = append(envelopes, envelope)
	}
	return envelopes, nil
}

func (repository EnvelopeRepositorio) Insert(envelope models.Envelope) (uint, error) {
	statement, erro := repository.sql.Prepare("Insert into envelopes(titulo,valor,observacao) values(?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(envelope.Titulo, envelope.Valor, envelope.Observacao)

	if erro != nil {
		return 0, erro
	}
	envelopeId, _ := result.LastInsertId()

	return uint(envelopeId), nil
}

func (repository EnvelopeRepositorio) GetEnvelopePorId(envelopeId uint) (models.Envelope, error) {
	linha, erro := repository.sql.Query("select id, titulo, valor, observacao from envelopes where id = ?", envelopeId)
	if erro != nil {
		return models.Envelope{}, erro
	}
	defer linha.Close()

	var envelope models.Envelope
	if linha.Next() {
		if erro = linha.Scan(
			&envelope.Id,
			&envelope.Titulo,
			&envelope.Valor,
			&envelope.Observacao,
		); erro != nil {
			return models.Envelope{}, erro
		}
	}

	return envelope, nil
}

func (repository EnvelopeRepositorio) DeleteById(envelopeId uint) error {
	statement, erro := repository.sql.Prepare("delete from envelopes where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(envelopeId)
	return erro
}

func (repository EnvelopeRepositorio) AtualizarEnvelope(envelope models.Envelope) error {
	statement, erro := repository.sql.Prepare("update envelopes set titulo = ?,valor = ? , observacao =? where id =?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(envelope.Titulo, envelope.Valor, envelope.Observacao, envelope.Id)
	return erro
}

func _inserirValorAModel(linha *sql.Rows, envelope *models.Envelope) error {
	if erro := linha.Scan(
		&envelope.Id,
		&envelope.Titulo,
		&envelope.Valor,
		&envelope.Observacao,
	); erro != nil {
		return erro
	}
	return nil
}
