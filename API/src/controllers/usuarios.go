package controllers

import (
	"api/src/authentication"
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = usuario.Preparar("cadastro"); err != nil {
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

	parametros := mux.Vars(r)

	id, err := strconv.ParseInt(parametros["usuarioId"], 10, 32)
	if err != nil {
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
	usuario, err := repositorio.BuscarPorID(uint64(id))
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)

}

// Atualiza um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	id, err := strconv.ParseInt(parametros["usuarioId"], 10, 32)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	usuarioIDNoToken, err := authentication.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioIDNoToken != uint64(id) {
		respostas.Erro(w, http.StatusForbidden, errors.New("Você não tem permissão para atualizar este usuário"))
		return
	}

	corpoDaRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(corpoDaRequisicao, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("edicao"); err != nil {
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
	err = repositorio.AtualizaUsuarioPorID(uint64(id), usuario)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)

}

// Deleta um usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	id, err := strconv.ParseInt(parametros["usuarioId"], 10, 32)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuarioIdNoToken, err := authentication.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioIdNoToken != uint64(id) {
		respostas.Erro(w, http.StatusForbidden, errors.New("Você não tem permissão para deletar este usuário"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	err = repositorio.DeletarUsuarioPorID(uint64(id))
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

// Deleta um usuário no banco de dados
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseInt(parametros["usuarioId"], 10, 32)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	seguidor_ID, err := authentication.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if seguidor_ID == uint64(usuarioID) {
		respostas.Erro(w, http.StatusForbidden, errors.New("Você náo pode seguir a si mesmo"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	err = repositorio.SeguirUsuario(uint64(usuarioID), uint64(seguidor_ID))
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func DeixarDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseInt(parametros["usuarioId"], 10, 32)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	seguidor_ID, err := authentication.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if seguidor_ID == uint64(usuarioID) {
		respostas.Erro(w, http.StatusForbidden, errors.New("Você náo pode deixar de seguir a si mesmo"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	err = repositorio.DeixarDeSeguirUsuario(uint64(usuarioID), uint64(seguidor_ID))
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseInt(parametros["usuarioId"], 10, 32)
	if err != nil {
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
	seguidores, err := repositorio.BuscarSeguidores(uint64(usuarioID))
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseInt(parametros["usuarioId"], 10, 32)
	if err != nil {
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
	seguindo, err := repositorio.BuscarSeguindo(uint64(usuarioID))
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, seguindo)
}
