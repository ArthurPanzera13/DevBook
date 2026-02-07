package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Retorna uma resposta em JSON para a requisiçào
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

// Retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: erro.Error(),
	})
}
