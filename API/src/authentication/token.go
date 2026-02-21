package authentication

import (
	"api/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Cria o token com as permissões do usuário
func CriarToken(usuarioID int64) (string, error) {
	permissao := jwt.MapClaims{}
	permissao["authorized"] = true
	permissao["exp"] = time.Now().Add(time.Hour * 1).Unix()
	permissao["usuarioID"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissao)

	return token.SignedString([]byte(config.SecretKey))

}

// Verifica se o token passado é válido
func ValidarToken(r *http.Request) error {

	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, retornaChaveDeVerificacao)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornaChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
