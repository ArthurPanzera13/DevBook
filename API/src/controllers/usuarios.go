package controllers

import "net/http"

// Insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Criando usuário"))

}

// Busca todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Buscando todos os usuários"))

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
