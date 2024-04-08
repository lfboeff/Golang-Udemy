package main

import "fmt"

func alunoEstaAprovado(n1, n2 float32) bool {

	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6")

}

func recuperarExecucao() {

	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}

}

func main() {
	fmt.Println(alunoEstaAprovado(9, 7))
	fmt.Println("Pós execução")
}

// antes de ocorrer o "panic", o programa executa tudo o que tiver sido "deferred"

// sem o "recover", a execução do programa será encerrada quando houver um "panic"
