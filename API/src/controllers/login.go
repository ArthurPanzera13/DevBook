package controllers

import (
	"api/src/authentication"
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Aitentica um usuario na API
func Login(w http.ResponseWriter, r *http.Request) {

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

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuarioEncontrado, err := repositorio.BuscarUsuarioPorEmail(usuario.Email)

	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(usuarioEncontrado)

	if err = security.VerificarSenha(usuario.Senha, usuarioEncontrado.Senha); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CriarToken(int64(usuarioEncontrado.ID))

	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
	}

	w.Write([]byte(token))
}
