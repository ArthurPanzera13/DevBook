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

// Deletar usuário por ID
func (repositorio usuarios) SeguirUsuario(usuario_id uint64, seguidor_id uint64) error {
	statemant, err := repositorio.db.Prepare("INSERT INTO SEGUIDORES (usuario_id, seguidor_id) VALUES (?, ?)")
	if err != nil {
		return err
	}

	defer statemant.Close()

	_, err = statemant.Exec(usuario_id, seguidor_id)
	if err != nil {
		return err
	}

	return nil
}

func (repositorio usuarios) DeixarDeSeguirUsuario(usuario_id uint64, seguidor_id uint64) error {
	statemant, err := repositorio.db.Prepare("DELETE FROM SEGUIDORES WHERE usuario_id = ? AND seguidor_id = ?")
	if err != nil {
		return err
	}

	defer statemant.Close()

	_, err = statemant.Exec(usuario_id, seguidor_id)
	if err != nil {
		return err
	}

	return nil
}

func (repositorio usuarios) BuscarSeguidores(usuario_id uint64) ([]models.Usuario, error) {
	statemant, err := repositorio.db.Prepare("SELECT u.id, u.nome, u.nickname, u.email FROM USUARIOS u INNER JOIN SEGUIDORES s ON u.id = s.seguidor_id WHERE s.usuario_id = ?")
	if err != nil {
		return nil, err
	}

	defer statemant.Close()

	linhas, err := statemant.Query(uint64(usuario_id))
	if err != nil {
		return nil, err
	}

	var seguidoresRetornados []models.Usuario
	for linhas.Next() {
		var seguidor models.Usuario
		if err = linhas.Scan(&seguidor.ID, &seguidor.Nome, &seguidor.Nick, &seguidor.Email); err != nil {
			return nil, err
		}

		seguidoresRetornados = append(seguidoresRetornados, seguidor)
	}

	return seguidoresRetornados, nil
}

func (repositorio usuarios) BuscarSeguindo(usuario_id uint64) ([]models.Usuario, error) {
	statemant, err := repositorio.db.Prepare("SELECT u.id, u.nome, u.nickname, u.email FROM USUARIOS u INNER JOIN SEGUIDORES s ON u.id = s.usuario_id WHERE s.seguidor_id = ?")
	if err != nil {
		return nil, err
	}

	defer statemant.Close()

	linhas, err := statemant.Query(uint64(usuario_id))
	if err != nil {
		return nil, err
	}

	var seguindoRetornados []models.Usuario
	for linhas.Next() {
		var seguindo models.Usuario
		if err = linhas.Scan(&seguindo.ID, &seguindo.Nome, &seguindo.Nick, &seguindo.Email); err != nil {
			return nil, err
		}

		seguindoRetornados = append(seguindoRetornados, seguindo)
	}

	return seguindoRetornados, nil
}

func (repositorio usuarios) BuscarSenhaUsuario(usuarioID uint64) (string, error) {
	var senha string

	err := repositorio.db.QueryRow(
		"SELECT senha FROM USUARIOS WHERE id = ?",
		usuarioID,
	).Scan(&senha)

	if err != nil {
		return "", err
	}

	return senha, nil
}

func (repositorio usuarios) AtualizarSenha(usuario_id uint64, novaSenha string) error {
	statemant, err := repositorio.db.Prepare("UPDATE USUARIOS SET senha = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer statemant.Close()

	_, err = statemant.Exec(novaSenha, usuario_id)
	if err != nil {
		return err
	}

	return nil
}
