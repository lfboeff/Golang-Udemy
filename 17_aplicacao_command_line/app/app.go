package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de CLI pronta para ser executada
func Gerar() *cli.App {

	app := cli.NewApp()

	app.Name = "Aplicação de Linha de Comando"

	app.Usage = "Busca IPs e Nomes de Servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIPs,
		},
		{
			Name:   "servidor",
			Usage:  "Busca Nomes de Servidores na internet",
			Flags:  flags,
			Action: buscaServidores,
		},
	}

	return app
}

func buscarIPs(c *cli.Context) {

	host := c.String("host")
	fmt.Println(host)

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscaServidores(c *cli.Context) {

	host := c.String("host")
	fmt.Println(host)

	servidores, err := net.LookupNS(host) // NS = "name server"
	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}
