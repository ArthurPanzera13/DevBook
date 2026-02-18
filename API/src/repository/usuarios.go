package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Consulta todos os usuários que atendem ao filtro de nome ou nickname
func (respositorio usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	statemant, err := respositorio.db.Prepare("SELECT id, nome, nickname, email FROM USUARIOS WHERE LOWER (nome) LIKE ? OR LOWER(nickname) LIKE ?")
	if err != nil {
		return nil, err
	}

	defer statemant.Close()

	linhas, err := statemant.Query(nomeOuNick, nomeOuNick)
	if err != nil {
		return nil, err
	}

	var usuariosRetornados []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario
		if err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email); err != nil {
			return nil, err
		}

		usuariosRetornados = append(usuariosRetornados, usuario)
	}
	return usuariosRetornados, nil
}

// Buscar usuário por ID
func (repositorio usuarios) BuscarPorID(id uint64) (models.Usuario, error) {
	statemant, err := repositorio.db.Prepare("SELECT id, nome, nickname, email FROM USUARIOS WHERE id = ?")
	if err != nil {
		return models.Usuario{}, err
	}

	defer statemant.Close()

	var usuario models.Usuario
	err = statemant.QueryRow(id).Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email)

	return usuario, nil
}

// Atualizar usuário por ID
func (repositorio usuarios) AtualizaUsuarioPorID(id uint64, usuario models.Usuario) error {
	statemant, err := repositorio.db.Prepare("UPDATE USUARIOS SET nome = ?, nickname = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer statemant.Close()

	_, err = statemant.Exec(usuario.Nome, usuario.Nick, usuario.Email, id)
	if err != nil {
		return err
	}

	if _, err = statemant.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); err != nil {
		return err
	}

	return nil
}

// Deletar usuário por ID
func (repositorio usuarios) DeletarUsuarioPorID(id uint64) error {
	statemant, err := repositorio.db.Prepare("DELETE FROM USUARIOS WHERE id = ?")
	if err != nil {
		return err
	}

	defer statemant.Close()

	_, err = statemant.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

// Busca usuário por email
func (repositorio usuarios) BuscarUsuarioPorEmail(email string) (models.Usuario, error) {
	linha, err := repositorio.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)

	if err != nil {
		return models.Usuario{}, err
	}

	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil

}
