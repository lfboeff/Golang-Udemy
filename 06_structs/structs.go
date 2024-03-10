package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {

	var u1 usuario
	fmt.Println(u1)

	u1.nome = "Felipe"
	u1.idade = 33
	fmt.Println(u1)

	u2 := usuario{nome: "Natalia", idade: 27, endereco: endereco{rua: "10 de Setembro", numero: 150}}
	fmt.Println(u2)

	u3 := usuario{nome: "Maria"}
	fmt.Println(u3.nome)

}
