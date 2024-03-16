package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println("Início do programa.")

	canal := multiplexador(escrever("Olá mundo!"), escrever(("Programando em Go")))

	for i := 0; i < 10; i++ {
		// fmt.Printf("Aguardando mensagem chegar no canal... %v\n", i)
		fmt.Println(<-canal)
		// fmt.Println("Mensagem recebida!")
	}

	// fmt.Println("Fim do programa.")
}

func multiplexador(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	// fmt.Println("Entrada em escrever()")
	canal := make(chan string)

	go func() {
		i := 1
		for {
			// fmt.Println("Entrada no for de escrever()")
			canal <- fmt.Sprintf("Valor recebido: %s - %v", texto, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
			// fmt.Println("Saída do for de escrever()")
			i++
		}
	}()
	// fmt.Println("Saída de escrever()")
	return canal
}
