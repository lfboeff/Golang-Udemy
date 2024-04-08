package main

// para criarmos o arquivo go.mod, usamos o comando "$ go mod init linha_de_comando"

import (
	"fmt"
	"linha_de_comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Início da aplicação:")

	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
