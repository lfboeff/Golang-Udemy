package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go escrever("Olá mundo!", &waitGroup)

	go escrever("Programando em Go", &waitGroup)

	fmt.Println("Aguardando o encerramento do programa.")

	waitGroup.Wait()

	fmt.Println("Programa encerrado!")

}

func escrever(texto string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
	waitGroup.Done()
}
