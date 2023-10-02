package repository

import (
	"api/src/models/envelope"
	"database/sql"
)

type EnvelopeRepositorio struct {
	sql *sql.DB
}

func NewInstanceEnvelope(sql *sql.DB) *EnvelopeRepositorio {
	return &EnvelopeRepositorio{sql}
}

func (repository EnvelopeRepositorio) GetEnvelopes(userId uint) ([]envelope.Envelope, error) {

	linhas, erro := repository.sql.Query("Select id, titulo, valor, observacao from envelopes where usuario_id = ?", userId)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var envelopes []envelope.Envelope
	for linhas.Next() {
		var envelope envelope.Envelope
		erro = _inserirValorAModel(linhas, &envelope)
		if erro == nil {
			envelopes = append(envelopes, envelope)
		}
	}
	return envelopes, erro
}

func (repository EnvelopeRepositorio) GetEnvelopePorNome(nome string, userId uint) ([]envelope.Envelope, error) {
	linhas, erro := repository.sql.Query("Select id, titulo, valor, observacao from envelopes where titulo LIKE ? and usuario_id = ?", "%"+nome+"%", userId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var envelopes []envelope.Envelope
	for linhas.Next() {
		var envelope envelope.Envelope
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

func (repository EnvelopeRepositorio) Insert(envelope envelope.Envelope) (uint, error) {
	statement, erro := repository.sql.Prepare("Insert into envelopes(titulo,valor,observacao,usuario_id) values(?,?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(envelope.Titulo, envelope.Valor, envelope.Observacao, envelope.Usuario)

	if erro != nil {
		return 0, erro
	}
	envelopeId, _ := result.LastInsertId()

	return uint(envelopeId), nil
}

func (repository EnvelopeRepositorio) GetEnvelopePorId(envelopeId, userId uint) (envelope.Envelope, error) {
	linha, erro := repository.sql.Query("select id, titulo, valor, observacao from envelopes where id = ? and usuario_id = ?", envelopeId, userId)
	if erro != nil {
		return envelope.Envelope{}, erro
	}
	defer linha.Close()

	var entity envelope.Envelope
	if linha.Next() {
		if erro = linha.Scan(
			&entity.Id,
			&entity.Titulo,
			&entity.Valor,
			&entity.Observacao,
		); erro != nil {
			return envelope.Envelope{}, erro
		}
	}

	return entity, nil
}

func (repository EnvelopeRepositorio) DeleteById(envelopeId, userId uint) error {
	statement, erro := repository.sql.Prepare("delete from envelopes where id = ? and usuario_id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(envelopeId, userId)
	return erro
}

func (repository EnvelopeRepositorio) AtualizarEnvelope(envelope envelope.Envelope, userId uint) error {
	statement, erro := repository.sql.Prepare("update envelopes set titulo=?, valor=? , observacao =? where id =? and usuario_id =?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(envelope.Titulo, envelope.Valor, envelope.Observacao, envelope.Id, userId)
	return erro
}

func _inserirValorAModel(linha *sql.Rows, envelope *envelope.Envelope) error {
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
