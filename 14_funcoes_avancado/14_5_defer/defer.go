package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func funcao3() {
	fmt.Println("Executando a função 3")
}

func alunoEstaAprovado(n1, n2 float32) bool {

	defer fmt.Printf("Média calculada. O resultado será retornado agora. Aluno aprovado: ")

	fmt.Println("Entrando na função para verificar se o aluno foi aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	} else {
		return false
	}

}

func main() {

	defer funcao1() // a clausula "defer" atrasa a execução de "funcao1()"
	defer funcao2()
	funcao3()

	fmt.Println("---")

	fmt.Println(alunoEstaAprovado(10, 10))

	fmt.Println("---")

}
