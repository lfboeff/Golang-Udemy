package main

import (
	"fmt"
	"time"
)

func main() {

	// Paralelismo != Concorrência
	// Paralelismo: tarefas executadas ao mesmo tempo em diferentes núcleos do processador
	// Concorrência: tarefas podem ou não estar sendo executadas ao mesmo tempo

	go escrever("Olá mundo!") // goroutine
	escrever("Programando em Go")

	// a goroutine acima faz com que o programa execute a função de baixo mesmo antes de ter concluído a execução da função de cima
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
