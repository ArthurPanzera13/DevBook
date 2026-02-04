package repository

import (
	"api/src/models"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario models.Usuario) (uint64, error) {

	statemant, err := repositorio.db.Prepare("INSERT INTO USUARIOS (nome, nickname, email, senha) VALUES(?,?,?,?)")

	if err != nil {
		return 0, err
	}
	defer statemant.Close()

	result, err := statemant.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoID), nil
}
