package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma senha (string) e retorna o hash dela
func Hash(senha string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

}

// VerificarSenha compara um hash e uma senha e retorna se os dois são iguais
func VerificarSenha(senhaComHash, senha string) error {

	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senha))

}
