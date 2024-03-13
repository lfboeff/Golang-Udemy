package app

import "github.com/urfave/cli"

// Gerar vai retornar a aplicação de CLI pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet"

	return app
}
