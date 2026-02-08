package security

import "golang.org/x/crypto/bcrypt"

// Hash gera um hash para a senha recebida
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara a senha recebida com o hash armazenado
func VerificarSenha(senhaString string, senhaHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
