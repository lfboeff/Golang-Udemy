package main

import "fmt"

func main() {

	usuario1 := map[string]string{ // as chaves de um map devem ser todas do mesmo tipo, assim como os valores devem ser todos do mesmo tipo
		"nome":      "Flávio",
		"sobrenome": "Dapper",
		"signo":     "Peixes",
	}
	fmt.Println(usuario1)
	fmt.Println(usuario1["nome"])

	fmt.Println()
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Felipe",
			"último":   "Dapper",
		},
		"local": {
			"rua":    "25 de Julho",
			"cidade": "São Leopoldo",
		},
		"curso": {
			"nome": "Engenharia",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"] + " " + usuario2["nome"]["último"] + " - " + usuario2["curso"]["nome"] + " - " + usuario2["local"]["cidade"])

	// Deletar uma chave usando a função nativa "delete":

	fmt.Println()
	delete(usuario1, "signo")
	fmt.Println(usuario1)
	delete(usuario2["local"], "rua")
	fmt.Println(usuario2)

	// Adicionar uma chave ou alterar uma já existente:

	fmt.Println()
	usuario1["hobby"] = "pescar"
	fmt.Println(usuario1)

	usuario2["hobby"] = map[string]string{
		"preferido": "Namorar a Natália",
		"outros":    "Programação",
	}
	fmt.Println(usuario2)

	usuario2["hobby"]["outros"] = "Programar com a Natália"
	fmt.Println(usuario2)

	usuario2["academia"] = map[string]string{"cidade": "DI"}
	fmt.Println(usuario2)
}
