module my_module

go 1.22.1

// podemos instalar bibliotecas externas a partir do comando "$ go get github.com/badoux/checkmail", por exemplo
require github.com/badoux/checkmail v1.2.4

// para remover qualquer dependência que não está (mais) ßsendo usada no projeto, podemos usar o comando "$ go mod tidy"
// o qual irá atualizar o arquivo go.mod para que ele contenha apenas as dependências necessárias para o projeto funcionar
