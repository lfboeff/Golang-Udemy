package modelos

import (
	"api_mod/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"` // "omitempty" faz com que dados nulos sejam ignorados pelo construtor json
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {

	if err := usuario.validar(etapa); err != nil {
		return err
	}

	if err := usuario.formatar(etapa); err != nil {
		return err
	}

	return nil

}

func (usuario *Usuario) validar(etapa string) error {

	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o email é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("o formato do email inserido não é válido")
	}

	if etapa == "cadastro" {
		if usuario.Senha == "" {
			return errors.New("a senha é obrigatória e não pode estar em branco")
		}
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {

	usuario.Nome = strings.TrimSpace(usuario.Nome)

	usuario.Nick = strings.TrimSpace(usuario.Nick)

	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {

		senhaComHash, err := seguranca.Hash(usuario.Senha)
		if err != nil {
			return err
		}
		usuario.Senha = string(senhaComHash)
	}

	return nil

}
