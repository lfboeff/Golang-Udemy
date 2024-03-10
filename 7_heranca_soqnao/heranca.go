package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	// se fizéssemos "pessoa pessoa", teremos que usar "estudante.pessoa.nome", mas o que queremos é poder usar "estudante.nome" simplesmente
	curso        string
	universidade string
}

func main() {

	p1 := pessoa{nome: "Daiana", sobrenome: "Cabeçuda", idade: 26, altura: 173}
	fmt.Println(p1)

	fmt.Println()
	e1 := estudante{pessoa: p1, curso: "Engenharia", universidade: "UFRGS"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.universidade)

}
