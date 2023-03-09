package usuario

import (
	"api/src/commons/seguranca"
	"errors"
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint   `json:"id"`
	Avatar   string `json:"avatar,omitempty"`
	Nome     string `json:"nome"`
	Username string `json:"username"`
	Senha    string
	Email    string `json:"email,omitempty"`
}

func (u *Usuario) Check() error {
	if erro := u.validateUsuario(); erro != nil {
		return erro
	}
	if erro := u.formataUsuario(); erro != nil {
		return erro
	}
	return nil
}

func (u *Usuario) formataUsuario() error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)

	senhaComHash, erro := seguranca.Hash(u.Senha)
	if erro != nil {
		return erro
	}

	u.Senha = string(senhaComHash)

	return nil
}

func (u *Usuario) validateUsuario() error {
	if erro := validated(u.Nome, "nome"); erro != nil {
		return erro
	}

	if erro := validated(u.Username, "username"); erro != nil {
		return erro
	}

	if erro := validateEmail(u.Email); erro != nil {
		return erro
	}

	if erro := validated(u.Senha, "senha"); erro != nil {
		return erro
	}
	return nil
}

func validated(valor, campo string) error {
	if valor == "" {
		return fmt.Errorf("o campo %s não pode ser vazio", valor)
	}
	return nil
}

func validateEmail(email string) error {
	if erro := validated(email, "email"); erro != nil {
		return erro
	}

	if erro := checkmail.ValidateFormat(email); erro != nil {
		return errors.New("email não é válido")
	}

	return nil
}
