package main

import "fmt"

func escrever() {
	fmt.Println("Escrevendo...")
}

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Usuario %v foi salvo.\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	(*u).idade++
}

func main() {

	usuario1 := usuario{"Felipe", 33}
	usuario1.salvar()

	fmt.Println("---")

	usuario2 := usuario{"Luis", 27}
	usuario2.salvar()
	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Printf("Usuário %v é maior de idade: %v\n", usuario2.nome, maiorDeIdade)
	usuario2.fazerAniversario()
	fmt.Printf("O usuário %v agora tem %v anos, pois fez aniversário!\n", usuario2.nome, usuario2.idade)

}
