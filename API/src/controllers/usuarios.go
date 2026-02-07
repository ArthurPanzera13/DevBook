package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadGateway, err)
		return
	}

	if err = usuario.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuario.ID, err = repositorio.Criar(usuario)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// Busca todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuarios, err := repositorio.Buscar(nomeOuNick)

	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)

}

// Busca um usuário específico no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Busca usuário"))

}

// Atualiza um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Atualizando usuário"))

}

// Deleta um usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Deletando usuário"))

}
