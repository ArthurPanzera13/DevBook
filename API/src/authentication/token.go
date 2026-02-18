package authentication

import (
	"api/config"
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
