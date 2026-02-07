package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nickname,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatório")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}

// Chama os métodos para formatar e validar o usuário
func (usuario *Usuario) Preparar(etapa string) error {

	if err := usuario.validar(etapa); err != nil {
		return err
	}

	usuario.formatar()

	return nil
}
