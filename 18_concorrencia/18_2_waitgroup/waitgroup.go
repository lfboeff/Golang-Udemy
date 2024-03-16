package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// vamos usar o waitgroup para sincronizar nossas goroutines, de forma que o programa só possa ser encerrado depois que elas estiverem encerradas também
	var waitGroup sync.WaitGroup

	// aqui estamos informando ao programa que 2 goroutines farão parte de waitGroup ("grupo de espera")
	waitGroup.Add(2)

	go func() { // função anônima da goroutine 1
		escrever("Olá mundo!")
		waitGroup.Done() // isso remove uma das goroutines do contador de waitGroup ("waitGroup = waitGroup - 1 = 2 - 1 = 1")
	}()

	go func() { // função anônima da goroutine 2
		escrever("Programando em Go")
		waitGroup.Done() // ("waitGroup = waitGroup - 1 = 1 - 1 = 0")
	}()

	fmt.Println("Aguardando o encerramento do programa.")

	waitGroup.Wait() // isso faz com que o programa efetivamente aguarde o contador de waitGroup chegar em 0 antes de encerrar

	fmt.Println("Programa encerrado!")

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
