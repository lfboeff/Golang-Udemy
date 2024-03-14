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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIPs,
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
